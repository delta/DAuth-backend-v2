package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	IsAuth(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
