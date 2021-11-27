package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
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

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	// query parameterの取得
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	database.DB.Delete(&product)

	return nil

}
