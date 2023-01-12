package salehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *saleHandler) CreateSale(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var sale entities.Sales
		p := presenter.NewSalePresenter()
		handle := handlers.NewAppHandler()

		fields := []string{"Order_id", "Employee_id", "Payment_method"}

		if err := c.BodyParser(&sale); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.SaleErrorResponse(err))
		}

		if err := handle.ValidateFields(fields, sale); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.SaleErrorResponse(err))
		}

		createdSale, err := ca.Sale.CreateSale(&sale)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.SaleErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.SaleSuccessResponse(createdSale))

	}
}
