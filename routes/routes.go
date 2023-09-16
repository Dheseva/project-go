package routes

import (
	"project-go/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())
	
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
	app.Get("/api/user/:id", controller.UpdateDataUser)

	app.Post("/api/registergun", controller.RegisterGun)
	app.Get("/api/gun/:id", controller.AirsoftgunView)
	app.Get("/api/gunep/:id", controller.AirsoftgunUpdate)
	app.Post("/api/guned/:id", controller.AirsoftgunDelete)

	app.Post("/api/foruse", controller.ForUse)
	app.Get("/api/foruse/:id", controller.ForUseView)
	app.Get("/api/forusep/:id", controller.ForUseUpdate)
	app.Post("/api/forused/:id", controller.ForUseDelete)
}