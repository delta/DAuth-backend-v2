package router

import (
	"github.com/delta/DAuth-backend-v2/middleware"
	"github.com/delta/DAuth-backend-v2/registry"
	"github.com/gofiber/fiber/v2"
)

func NewAppRouter(app *fiber.App, registry registry.Registry) {
	api := app.Group("/api", middleware.Validate())
	appController := registry.NewAppController()

	NewAuthRouter(api.Group("/auth"), appController.Auth)
}
