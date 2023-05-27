package router

import (
	"github.com/gofiber/fiber/v2"
)


type AuthRouter interface {
	Route(app *fiber.App)
}
