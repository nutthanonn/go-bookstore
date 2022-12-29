package customertest

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/fiber/v2/utils"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
	"github.com/nutthanonn/go-clean-architecture/api/routers"
	"github.com/nutthanonn/go-clean-architecture/pkg/registry"
)

func Test_CreateCustomer(t *testing.T) {
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
	routers.CustomerHandler(api, r.NewAppController())

	t.Run("FAIL CREATE CUSTOMER", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/customer", nil)
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)

		utils.AssertEqual(t, nil, err, "app.Test(req)")
		utils.AssertEqual(t, fiber.StatusBadRequest, res.StatusCode, "Status code")
	})
}
