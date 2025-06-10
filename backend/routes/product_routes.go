package routes

import (
	"github.com/henriqueAfonso/inventory-system/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/products", controllers.GetProducts)
	api.Post("/products", controllers.CreateProduct)
}
