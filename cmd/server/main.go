// cmd/server/main.go
package main

import (
	"github.com/Jere283/autek_rest_api/config"
	"github.com/Jere283/autek_rest_api/database"
	"github.com/Jere283/autek_rest_api/handlers"
	"github.com/Jere283/autek_rest_api/repository"
	"github.com/Jere283/autek_rest_api/routes"
	"github.com/Jere283/autek_rest_api/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to Database
	database.ConnectDB()

	// Initialize app
	app := fiber.New()

	// Dependency Injection
	userRepo := repository.UserRepository{}
	userService := service.UserService{UserRepo: userRepo}
	userHandler := handlers.UserHandler{UserService: userService}

	// Setup routes
	routes.SetupRoutes(app, &userHandler)

	// Start the server
	app.Listen(":3000")
}
