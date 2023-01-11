package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/nutthanonn/go-clean-architecture/api/config"
)

func Limiter() interface{} {
	return limiter.New(config.NewLimiterConfig())
}
