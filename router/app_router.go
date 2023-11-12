package router

import (
	// Include docs for Swagger UI generation
	// _ "github.com/delta/DAuth-backend-v2/docs"
	"github.com/delta/DAuth-backend-v2/middleware"
	"github.com/delta/DAuth-backend-v2/registry"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewAppRouter(app *fiber.App, registry registry.Registry) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api", middleware.Validate())
	appController := registry.NewAppController()

	NewAuthRouter(api.Group("/auth"), appController.Auth)
}
