package handlers

import (
	"go-api-fiber-gorm/database"
	"go-api-fiber-gorm/models/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUserHandler(ctx *fiber.Ctx) error {

	var users []entity.User

	result := database.DB.Debug().Find(&users)
	log.Println("result", result)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)

}
