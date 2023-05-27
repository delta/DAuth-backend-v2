package main

import (
	"fmt"

	config "github.com/delta/DAuth-backend-v2/config/impl"
	controller "github.com/delta/DAuth-backend-v2/controller/impl"
	"github.com/delta/DAuth-backend-v2/database"
	"github.com/delta/DAuth-backend-v2/entity"
	repository "github.com/delta/DAuth-backend-v2/repository/impl"
	router "github.com/delta/DAuth-backend-v2/router/impl"
	service "github.com/delta/DAuth-backend-v2/service/impl"
	"github.com/fatih/color"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.New()
	database := database.Connect(config)

	entity.Migrate(database)

	resourceRepository := repository.NewResourceOwnerRepositoryImpl(database)
	authService := service.NewAuthServiceImpl(resourceRepository)
	authController := controller.NewAuthControllerImpl(authService)
	authRouter := router.NewAuthRouterImpl(authController)

	app := fiber.New()
	authRouter.Route(app)
	
	err := app.Listen(":" + config.Get("PORT"))

	if err != nil {
		fmt.Println(color.RedString("Error starting the server: %s", err))
	}
}
