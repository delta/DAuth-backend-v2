package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Validate() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// to pass the middleware
		// validation logic goes here
		if true {
			err := c.Next()
			if err != nil {
				return nil
			}
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("Access denied")
		}

		return nil
	}
}
