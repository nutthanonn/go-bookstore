package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func CSRFCookie() interface{} {

	return csrf.New(csrf.Config{
		KeyLookup:      "form:csrf_token",
		CookieName:     "csrf_cookie",
		CookieHTTPOnly: true,
	})
}
