package main

import (
	"github.com/henriqueAfonso/inventory-system/backend/config"
	"github.com/henriqueAfonso/inventory-system/backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
