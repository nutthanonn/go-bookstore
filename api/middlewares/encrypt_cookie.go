package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/nutthanonn/go-clean-architecture/api/config"
)

func EncryptCookie() interface{} {
	return encryptcookie.New(config.NewEncryptCookieConfig())
}
