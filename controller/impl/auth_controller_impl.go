package impl

import (
	"fmt"

	config "github.com/delta/DAuth-backend-v2/config/impl"
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

func (impl *authControllerImpl) Login(c *fiber.Ctx) error {

	var req model.LoginRequest

	err := c.BodyParser(&req)

	if err != nil {
		fmt.Printf("Error parsing: %s", err)
		return err
	}

	var userEmail entity.Email

	if userEmail, err = impl.emailService.FindByEmail(c.Context(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	var userDetails entity.ResourceOwner

	if userDetails, err = impl.resourceService.FindByEmailID(c.Context(), userEmail.ID); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	logger := utils.GetControllerLogger("Login")
	logger.Info(fmt.Sprintf("UserID:%v,Route:", userDetails.ID) + "Login")

	if utils.CheckPasswordHash(req.Password, userDetails.Password) {

		jwtToken, err := utils.GenerateToken(userDetails.ID)
		if err != nil {
			fmt.Println(err)
		}

		return c.Status(fiber.StatusOK).JSON(jwtToken)
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Invalid Credentials")
}

func (impl *authControllerImpl) IsAuth(c *fiber.Ctx) error {

	config := config.New()
	accessTokenName := config.Get("ACCESS_TOKEN_NAME")
	userToken := c.Cookies(accessTokenName)

	userID, err := utils.VerifyToken(userToken)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	logger := utils.GetControllerLogger("Is-Auth")
	logger.Info(fmt.Sprintf("UserID:%v,Route:", userID) + "Is-Auth")

	if userID != 0 {
		userDetails, err := impl.resourceService.FindByID(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(userDetails)
	}

	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}

func (impl *authControllerImpl) EditProfile(c *fiber.Ctx) error {
	// var req model.EditProfile

	// if ok := utils.IsValidPhoneNumber(req.PhoneNumber); !ok {

	// }
	return c.Status(fiber.StatusOK).SendString("Profile Updated")
}
