package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	// OrderItemsをPreloadすることで、orderテーブルのorderItemsレコードを
	// 別テーブルであるOrderItemsから事前に取得できる
	database.DB.Preload("OrderItems").Find(&orders)

	// first name と last name を組み合わせてfull name を返すようにする
	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}

	return c.JSON(orders)
}
