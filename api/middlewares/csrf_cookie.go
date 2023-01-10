package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/utils"
)

func CSRFCookie() interface{} {

	return csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	})
}
