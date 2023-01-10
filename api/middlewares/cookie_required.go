package middlewares

import "github.com/gofiber/fiber/v2"

func RequiredCookie() interface{} {
	return func(c *fiber.Ctx) error {
		if c.Cookies("access_token") == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false,
				"error":  "Unauthorized",
			})
		}

		c.Next()
		return nil
	}
}
