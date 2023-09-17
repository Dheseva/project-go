package database

import (
	"fmt"
	"os"
	"project-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	mysqlDb := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", 
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
	connection, err := gorm.Open(mysql.Open(mysqlDb), &gorm.Config{})

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