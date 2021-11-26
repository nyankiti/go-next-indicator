package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/midddlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

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
}
