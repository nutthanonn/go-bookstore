package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/nutthanonn/go-clean-architecture/api/config"
)

func CSRFCookie() interface{} {

	return csrf.New(config.NewCSRFConfig())
}
