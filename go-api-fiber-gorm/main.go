package main

import (
	"go-api-fiber-gorm/database"
	"go-api-fiber-gorm/database/migration"
	"go-api-fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// initiate DB
	database.DatabaseInit()

	// auto Migrate DB
	migration.RunMigration()

	// initiate Fiber
	app := fiber.New()

	// routes
	routes.RouteInit(app)

	app.Listen(":7000")
}
