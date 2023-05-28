package router

import (
	"github.com/delta/DAuth-backend-v2/controller"
	"github.com/gofiber/fiber/v2"
)


func NewAuthRouter(app fiber.Router,controller controller.AuthController){
	app.Post("/add", controller.Insert)
	app.Delete("/delete", controller.Remove)
}
