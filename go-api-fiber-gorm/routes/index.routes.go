package routes

import (
	"go-api-fiber-gorm/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	api.Get("/", handlers.WelcomeHandler)
	v1 := api.Group("/v1")
	user := v1.Group("/user")


	user.Get("/lists", handlers.GetAllUserHandler)
	user.Post("/add", handlers.AddUserHandler)
}
