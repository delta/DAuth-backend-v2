package impl

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/middleware"
	"github.com/delta/DAuth-backend-v2/router"
	"github.com/gofiber/fiber/v2"
)

type authRouterImpl struct {
	authController controller.AuthController
}

func NewAuthRouterImpl(authController controller.AuthController) router.AuthRouter {
	return &authRouterImpl{authController}
}

func (impl *authRouterImpl) Route(app *fiber.App) {
	app.Post("/auth/add", impl.authController.Insert)
	app.Delete("/auth/delete", middleware.Validate(), impl.authController.Remove)
}
