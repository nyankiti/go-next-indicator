package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/gofiber/fiber/v2"
	"sort"
	"strconv"
	"strings"
)

func Product(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	go database.ClearCache("products_frontend", "products_backend")

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	// query parameterの取得
	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)

	// idだけ決められたproduct インスタンスのポインターを渡すことで、そのidのデータを探し出し、
	//そのポインターにデータを格納してくれる
	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	// query parameterの取得
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)
	// Modelの引数に、どのModel(table)を使うかを示すためのものなので、以下のように内容のないmodels.Product structを渡しても同様に動作する
	//database.DB.Model(models.Product{}).Updates(&product)

	go database.ClearCache("products_frontend", "products_backend")

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	// query parameterの取得
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	database.DB.Delete(&product)

	go database.ClearCache("products_frontend", "products_backend")

	return nil
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	//var ctx = context.Background()
	//
	////まずcacheにデータがあるか確認する
	//result, err := database.Cache.Get(ctx, "products_frontend").Result()
	//
	//if err != nil {
	//	database.DB.Find(&products)
	//	// redisにcacheを登録する際はbyteデータである必要がある
	//	bytes, err := json.Marshal(products)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
	//		panic(errKey)
	//	}
	//}else{
	//json.Unmarshal([]byte(result), &products)
	//}

	// 上のredisによるcacheの処理がエラーするので一時的に以下のようにキャッシュなしでやる
	database.DB.Find(&products)
	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx) error {
	var products []models.Product
	//var ctx = context.Background()
	//
	////まずcacheにデータがあるか確認する
	//result, err := database.Cache.Get(ctx, "products_backend").Result()
	//
	//if err != nil {
	//	database.DB.Find(&products)
	//	// redisにcacheを登録する際はbyteデータである必要がある
	//	bytes, err := json.Marshal(products)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute)
	//}else{
	//json.Unmarshal([]byte(result), &products)
	//}

	// 上のredisによるcacheの処理がエラーするので一時的に以下のようにキャッシュなしでやる
	database.DB.Find(&products)

	var searchedProducts []models.Product

	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			// query paramsの値がtitleかdescriptionに含まれている物を返す
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}

	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			// 並び替えはsortを使って以下のように実現する
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}

	var total = len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	// 一つのpageに表示するアイテム数
	perPage := 9

	var data []models.Product

	// 2ページ目以降でかつ、そのページを全て埋めるProduct数がない時
	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchedProducts[(page-1)*perPage : total]
		// 通常のペジネーション
	} else if total >= page*perPage {
		data = searchedProducts[(page-1)*perPage : page*perPage]
		// データがないのに、大きなページにアクセスされた時など、その他の場合
	} else {
		data = []models.Product{}
	}

	//return c.JSON(data[(page-1)*perPage : page*perPage])
	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
