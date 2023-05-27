package impl

import (
	"fmt"

	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
	"github.com/delta/DAuth-backend-v2/service"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	authService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) controller.AuthController {
	return &authControllerImpl{authService}
}

func (impl *authControllerImpl) Insert(c *fiber.Ctx) error {
	var body model.Resource
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Printf("Error parsing: %s", err)
		return err
	}

	resource := entity.ResourceOwner{
		Name: body.Name,
	}

	response := impl.authService.Create(c.Context(), resource)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (impl *authControllerImpl) Remove(c *fiber.Ctx) error {
	var body model.Resource
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Printf("Error parsing: %s", err)
		return err
	}

	resource := entity.ResourceOwner{
		Name: body.Name,
	}

	impl.authService.Delete(c.Context(), resource)
	return c.Status(fiber.StatusAccepted).SendString("Success")
}
