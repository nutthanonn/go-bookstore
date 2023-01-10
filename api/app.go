package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/routers"
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
	routers.BookRouter(api, r.NewAppController())
	routers.EmployeeRouter(api, r.NewAppController())
	routers.CustomerRouter(api, r.NewAppController())
	routers.OrderRouter(api, r.NewAppController())
	routers.AuthRouter(api, r.NewAppController())

	app.Listen(":3000")
}
