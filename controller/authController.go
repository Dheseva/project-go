package controller

import (
	"project-go/database"
	"project-go/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name: data["name"],
		Password: password,
		Email: data["email"],
	}

	database.DB.Create(&user)
	return c.JSON(user)
}