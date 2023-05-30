package routes

import (
	"github.com/RianIhsan/ApiGoJwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
