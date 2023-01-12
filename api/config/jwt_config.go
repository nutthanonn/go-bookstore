package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
)

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": false,
		"error":  "Unauthorized",
	})

	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func NewJwtConfig() jwtware.Config {
	datastore.LoadEnv()

	TOKEN_SECRET_KEY := os.Getenv("TOKEN_SECRET_KEY")
	return jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(TOKEN_SECRET_KEY),
		SigningMethod:  "HS256",
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	}
}
