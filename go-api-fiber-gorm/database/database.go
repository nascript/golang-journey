package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/go-fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connet to MYSQL database")
	}

	fmt.Println("Connected to MYSQL database")
}
