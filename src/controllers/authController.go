package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// fiber contextのbodyParserによってdataのポインタにrequest bodyを格納する
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}



	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	// userが見つからない時
	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials User",
		})
	}

	// passwordが間違っている時
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials Password",
		})
	}

	payload := jwt.StandardClaims{
		// Itoa: Integer to ASCII
		Subject: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials new jwt",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour*24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// cookeiに格納したjwtをparseすることによってjwtの内容であるuser idを取得できる
	payload := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", payload.Subject).First(&user)

	return c.JSON(user)
}