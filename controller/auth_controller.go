package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Insert(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}
