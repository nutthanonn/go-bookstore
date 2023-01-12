package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nutthanonn/go-clean-architecture/api/config"
)

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(config.NewJwtConfig())
}
