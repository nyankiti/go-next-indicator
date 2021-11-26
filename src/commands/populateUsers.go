// 独自で実行するコマンドファイルなので、package main とする必要がある
package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
)

// このコマンドは　docker-compose exec backend sh  としてコンテナに入ってから実行する必要がある
// go run src/commands/populateUsers.go とすると実行できる

func main() {
	//独立したコマンドであるため、データベースへの接続を独自で行う必要がある
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
			Email: faker.Email(),
			IsAmbassador: true,

		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}
}
