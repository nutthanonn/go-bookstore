package datastore

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/goCleanArchitecture?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{})

	return db
}
