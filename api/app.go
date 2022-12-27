package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
	"github.com/nutthanonn/go-clean-architecture/api/routers"
	"github.com/nutthanonn/go-clean-architecture/pkg/registry"
)

func main() {
	db := datastore.NewDB()
	defer func() {
		sqlDB, err := db.DB()

		if err != nil {
			panic(err)
		}

		sqlDB.Close()
	}()

	app := fiber.New()
	api := app.Group("/api")
	r := registry.NewRegistry(db)
	routers.UserRouter(api, r.NewAppController())
	routers.BookRouter(api, r.NewAppController())

	app.Listen(":3000")
}
