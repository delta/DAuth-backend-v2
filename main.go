package main

import (
	"fmt"

	config "github.com/delta/DAuth-backend-v2/config/impl"
	"github.com/delta/DAuth-backend-v2/database"
	"github.com/delta/DAuth-backend-v2/registry"
	"github.com/delta/DAuth-backend-v2/router"
	"github.com/delta/DAuth-backend-v2/utils"
	"github.com/fatih/color"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.New()
	utils.InitLogger()
	db := database.Connect(config)
	database.Migrate(db)

	registry := registry.NewRegistry(db)

	app := fiber.New()

	router.NewAppRouter(app, registry)

	err := app.Listen(":" + config.Get("PORT"))

	if err != nil {
		fmt.Println(color.RedString("Error starting the server: %s", err))
	}
}
