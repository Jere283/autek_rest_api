// cmd/server/main.go
package main

import (
	"github.com/Jere283/autek_rest_api/config"
	"github.com/Jere283/autek_rest_api/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to Database
	database.ConnectDB()

	// Initialize app
	app := fiber.New()

	// Start the server
	app.Listen(":3000")
}
