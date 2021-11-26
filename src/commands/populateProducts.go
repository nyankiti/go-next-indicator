package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// このコマンドは　docker-compose exec backend sh  としてコンテナに入ってから実行する必要がある
// go run src/commands/populateProducts.go とすると実行できる

func main() {
	//独立したコマンドであるため、データベースへの接続を独自で行う必要がある
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.Product{
			Title: faker.Username(),
			Description: faker.Username(),
			Image: faker.URL(),
			// 10 ~ 100のうち、ランダムの数字を返すため、10を分けて足している
			Price: float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&ambassador)
	}
}
