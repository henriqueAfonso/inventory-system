package config

import (
	"fmt"

	"github.com/henriqueAfonso/inventory-system/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=inventory port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.Product{})
	DB = db
	fmt.Println("Database connected!")
}
