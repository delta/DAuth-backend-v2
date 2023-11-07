package impl

import (
	"fmt"

	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
	"github.com/delta/DAuth-backend-v2/service"
	"github.com/delta/DAuth-backend-v2/utils"

	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	authService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) controller.AuthController {
	return &authControllerImpl{authService}
}

// Insert inserts a new resource owner.
// @Summary Insert a new resource owner
// @Description Inserts a new resource owner.
// @Tags Auth
// @Accept json
// @Produce json
// @Param resource body model.Resource true "Resource object to be inserted"
// @Success 201 {object} model.Resource
// @Error 400 {string} BadRequest "Error parsing: {{errMsg}}"
// @Router /api/auth/add [post]
func (impl *authControllerImpl) Insert(c *fiber.Ctx) error {
	var body model.Resource
	err := c.BodyParser(&body)

	// Using Logger
	logger := utils.GetControllerLogger("Insert")
	logger.Info("Insert Controller Used")

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

// Remove removes a resource owner.
// @Summary Remove a resource owner
// @Description Removes a resource owner
// @Tags Auth
// @Accept json
// @Produce plain
// @Param resource body model.Resource true "Resource object to be removed"
// @Success 202 {string} string "Success"
// @Error 400 {string} BadRequest "Error parsing: {{errMsg}}"
// @Router /api/auth/delete [delete]
func (impl *authControllerImpl) Remove(c *fiber.Ctx) error {
	var body model.Resource
	err := c.BodyParser(&body)

	// Using Logger
	logger := utils.GetControllerLogger("Delete")
	logger.Info("Delete Controller Used")

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
