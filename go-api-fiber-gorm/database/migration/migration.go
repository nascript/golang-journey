package migration

import (
	"fmt"
	"go-api-fiber-gorm/database"
	"go-api-fiber-gorm/models/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
