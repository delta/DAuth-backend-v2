package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Insert(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
	IsAuth(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	EditProfile(c *fiber.Ctx) error
	// VerifyEmail(c *fiber.Ctx) error
}
