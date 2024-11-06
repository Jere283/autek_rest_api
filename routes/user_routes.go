package routes

import (
	"github.com/Jere283/autek_rest_api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")
	api.Get("/users/:id", userHandler.GetUser)
}
