package handlers

import (
	"go-api-fiber-gorm/database"
	"go-api-fiber-gorm/models/entity"
	"go-api-fiber-gorm/models/request"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddUserHandler(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}

	errCreateUser := database.DB.Debug().Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors":  http.StatusBadRequest,
			"message": "filed to store data",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"errors":  http.StatusOK,
		"message": "User Created",
		"data":    newUser,
	})
}
