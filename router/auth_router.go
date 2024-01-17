package router

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/registry"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRouter(app fiber.Router, controller controller.AuthController) {
	app.Get("/is-auth", controller.IsAuth)
	app.Post("/login", controller.Login)
}
func newAppRouter(app *fiber.App, registry *registry.Registry) {
	// Other routes...

	app.Post("/register/student", controller.RegisterStudent)
}
