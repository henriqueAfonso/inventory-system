package controllers

import (
	"github.com/henriqueAfonso/inventory-system/backend/config"
	"github.com/henriqueAfonso/inventory-system/backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	config.DB.Find(&products)
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	config.DB.Create(&product)
	return c.JSON(product)
}
