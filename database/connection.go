package database

import (
	"project-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	connection, err := gorm.Open(mysql.Open("root:@/go-receipt"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}