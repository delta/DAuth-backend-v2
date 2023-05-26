package main

import (
	// config "github.com/delta/DAuth-backend-v2/config/impl"
	// "github.com/delta/DAuth-backend-v2/database"
	// repository "github.com/delta/DAuth-backend-v2/repository/impl"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {


        return c.SendString("Hello, World!")
    })
    
    // config := config.New()
    // database := database.NewDatabase(config)


    app.Listen(":3000")
}
