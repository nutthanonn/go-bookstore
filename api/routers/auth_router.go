package routers

import (
	"github.com/gofiber/fiber/v2"
	authHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/auth_handler"
	"github.com/nutthanonn/go-clean-architecture/api/middlewares"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func AuthRouter(app fiber.Router, ca controller.AppController) {
	ah := authHandler.NewAuthHandler()

	app.Use(middlewares.EncryptCookie())
	app.Post("/signin", ah.SignIn(ca))

	app.Use(middlewares.RequiredCookie())
	app.Use(middlewares.CSRFCookie())

	app.Get("/cookie", func(c *fiber.Ctx) error {
		return c.SendString(c.Cookies("access_token"))
	})

	// handle csrf cookie
	app.Get("/cookie", func(c *fiber.Ctx) error {
		return c.SendString(c.Cookies("csrf_token"))
	})
}
