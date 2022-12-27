package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func CreateUser(co controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user entities.User
		p := presenter.NewUserPresenter()

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.UserErrorResponse(err))
		}

		createdUser, err := co.User.CreateUser(&user)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(p.UserSuccessResponse(createdUser))
	}
}

func GetUsers(co controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := co.User.ReadUser()
		p := presenter.NewUserPresenter()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.UserErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.UsersSuccessResponse(users))
	}
}

func DeleteUser(co controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")
		p := presenter.NewUserPresenter()

		if err := co.User.DeleteUser(ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.UserErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.UserDeleteSuccessResponse())
	}
}

func UpdateUser(co controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user entities.User
		p := presenter.NewUserPresenter()

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.UserErrorResponse(err))
		}

		updatedUser, err := co.User.UpdateUser(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.UserErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.UserSuccessResponse(updatedUser))
	}
}
