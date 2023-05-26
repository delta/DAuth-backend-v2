package impl

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/service"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct{
	authService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) controller.AuthController{
	return &authControllerImpl{authService}
}

func (impl *authControllerImpl) Login(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
