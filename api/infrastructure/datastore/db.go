package datastore

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&entities.User{},
		&entities.Employees{},
		&entities.Books{},
		// &entities.Customers{},
		// &entities.Orders{},
		// &entities.OrderDetails{},
		// &entities.Inventories{},
		// &entities.Sales{},
	); err != nil {
		panic(err)
	}

	return db
}
