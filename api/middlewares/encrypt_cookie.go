package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
)

func EncryptCookie() interface{} {
	datastore.LoadEnv()

	COOKIE_SECRET_KEY := os.Getenv("COOKIE_SECRET_KEY")

	return encryptcookie.New(encryptcookie.Config{
		Key:    COOKIE_SECRET_KEY,
		Except: []string{"csrf_"}, // exclude CSRF cookie
	})
}
