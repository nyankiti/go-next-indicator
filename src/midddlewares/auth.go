package midddlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// middlewareで動作しているので、次のrequestに返す意味で以下のようにreturnする
	return c.Next()
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	// uint 型はunsigned integer の略で、符号を気にする必要のない数字型としてよく使う
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		// errの場合は0をuser idとして返す
		return 0, err
	}

	// cookeiに格納したjwtをparseすることによってjwtの内容であるuser idを取得できる
	payload := token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}
