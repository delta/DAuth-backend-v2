package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Validate() func(c *fiber.Ctx) error {
	return func (c *fiber.Ctx) error {
		// to pass the middleware
		// validation logic goes here
		if (true) {
			c.Next()
		} else {
			c.Status(fiber.StatusBadRequest).SendString("Access denied")
		}

		return nil
	}
}
