package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/midddlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	//user admin--------------------------------------------------------------
	admin := api.Group("admin")
	// 認証前エンドポイント
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)
	// 認証後エンドポイント
	adminAuthenticated := admin.Use(midddlewares.IsAuthenticated)

	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)

	adminAuthenticated.Get("ambassador", controllers.Ambassador)
	adminAuthenticated.Get("products", controllers.Product)
	adminAuthenticated.Post("products", controllers.CreateProducts)
	// query parameterは以下のようにして渡す
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)

	adminAuthenticated.Get("users/:id/links", controllers.Link)

	adminAuthenticated.Get("orders", controllers.Orders)

	//ambassador------------------------------------------------------------
	ambassador := api.Group("ambassador")
	//認証前エンドポイント
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)
	ambassador.Get("products/frontend", controllers.ProductsFrontend)
	ambassador.Get("products/backend", controllers.ProductsBackend)
	// 認証後エンドポイント
	ambassadorAuthenticated := ambassador.Use(midddlewares.IsAuthenticated)

	ambassadorAuthenticated.Get("user", controllers.User)
	ambassadorAuthenticated.Post("logout", controllers.Logout)
	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
}
