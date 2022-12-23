package handlers

import "github.com/gofiber/fiber/v2"

func WelcomeHandler(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"Welcome": "GO API FIBER GORM by nascript",
	})

}
