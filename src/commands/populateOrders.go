package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// このコマンドは　docker-compose exec backend sh  としてコンテナに入ってから実行する必要がある
// go run src/commands/populateOrders.go とすると実行できる

func main() {
	//独立したコマンドであるため、データベースへの接続を独自で行う必要がある
	database.Connect()

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          qty,
				AdminRevenue:      0.9 * price * float64(qty),
				AmbassadorRevenue: 0.1 * price * float64(qty),
			})
		}

		// 以下のようにOrderテーブルに追加するメソッドを実行するだけで、
		// 外部キー関係にあるorderItemsテーブルも自動的に反映される
		database.DB.Create(&models.Order{
			// 1~31のisを持つuserをfakerで作成しているので、そのuserの誰かにランダムで紐付ける
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Complete:        true,
			OrderItems:      orderItems,
		})
	}
}
