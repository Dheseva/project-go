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
	connection.AutoMigrate(&models.UserData{})
	
	connection.AutoMigrate(&models.ForUse{})
	
	connection.AutoMigrate(&models.Airsoftgun{})
	connection.AutoMigrate(&models.AirsoftgunData{})
	
	
}

func AutoMigration(){
}