package impl

import (
	"fmt"
	"regexp"

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
		return c.Status(fiber.StatusBadRequest).SendString(ErrInvalidRequest.Error())
	}

	rex, err := regexp.MatchString(`^[\w-\.]+@nitt.edu`, req.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(ErrInvalidRequest.Error())
	}

	if !rex {
		return c.Status(fiber.StatusBadRequest).SendString(ErrInvalidRequest.Error())
	}

	logger := utils.GetControllerLogger(c.Path())

	var userEmail entity.Email

	if userEmail, err = impl.emailService.FindByEmail(c.Context(), req); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
	}

	var userDetails entity.ResourceOwner

	if userDetails, err = impl.resourceService.FindByEmailID(c.Context(), userEmail.ID); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
	}

	if utils.CheckPasswordHash(req.Password, userDetails.Password) {

		jwtToken, err := utils.GenerateToken(userDetails.ID)
		if err != nil {
			logger.Error(err.Error() + fmt.Sprintf("UserID:%v", userDetails.ID))
			return c.Status(fiber.StatusInternalServerError).SendString(ErrInternalServerError.Error())
		}

		logger.Info(fmt.Sprintf("UserID:%v", userDetails.ID))
		return c.Status(fiber.StatusOK).JSON(jwtToken)
	}
	return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
}

func (impl *authControllerImpl) IsAuth(c *fiber.Ctx) error {

	config := config.New()
	accessTokenName := config.Get("ACCESS_TOKEN_NAME")
	userToken := c.Cookies(accessTokenName)
	logger := utils.GetControllerLogger(c.Path())

	userID, err := utils.VerifyToken(userToken)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
	}

	if userID != 0 {
		userDetails, err := impl.resourceService.FindByID(c.Context(), userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
		}
		logger.Info(fmt.Sprintf("UserID:%v", userID))
		return c.Status(fiber.StatusOK).JSON(userDetails)
	}

	return c.Status(fiber.StatusUnauthorized).SendString(ErrUserNotFound.Error())
}
