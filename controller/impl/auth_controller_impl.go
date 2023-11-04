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
	resourceService service.ResourceService
	emailService    service.EmailService
}

func NewAuthControllerImpl(resourceService service.ResourceService, emailService service.EmailService) controller.AuthController {
	return &authControllerImpl{resourceService, emailService}
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

	response := impl.resourceService.Create(c.Context(), resource)
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

	impl.resourceService.Delete(c.Context(), resource)
	return c.Status(fiber.StatusAccepted).SendString("Success")
}

func (impl *authControllerImpl) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	err := c.BodyParser(&req)

	if err != nil {
		fmt.Printf("Error parsing: %s", err)
		return err
	}

	var userEmail entity.Email

	if userEmail = impl.emailService.FindByEmail(c.Context(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Email Not Found")
	}

	var userDetails entity.ResourceOwner

	if userDetails = impl.resourceService.FindByEmailID(c.Context(), userEmail.ID); userDetails == (entity.ResourceOwner{}) {
		return c.Status(fiber.StatusUnauthorized).SendString("User Not Found")
	}

	if utils.CheckPasswordHash(req.Password, userDetails.Password) {

		jwtToken, err := utils.GenerateToken(userDetails.ID)
		if err != nil {
			fmt.Println(err)
		}

		return c.Status(fiber.StatusAccepted).JSON(jwtToken)
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Invalid Credentials")
}

func (impl *authControllerImpl) IsAuth(c *fiber.Ctx) error {

	userToken := c.Cookies("access_token")

	userID, err := utils.VerifyToken(userToken)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	if userID != 0 {
		userDetails := impl.resourceService.FindByID(c.Context(), userID)
		return c.Status(fiber.StatusAccepted).JSON(userDetails)
	}

	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}

func (impl *authControllerImpl) EditProfile(c *fiber.Ctx) error {
	// var req model.EditProfile

	// if ok := utils.IsValidPhoneNumber(req.PhoneNumber); !ok {

	// }
	return c.Status(fiber.StatusUnauthorized).SendString("Profile Updated")
}
