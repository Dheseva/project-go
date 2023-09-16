package controller

import (
	"project-go/database"
	"project-go/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ForUse(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}

	foruse := models.ForUse{
		Use_for: data["use_for"],
		User_id: 2,
		Airsoftgun_id: 1,
		Comment: data["comment"],
		Expired_at: time.Now().Add(time.Hour * 360).UnixMilli(),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	database.DB.Create(&foruse)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success register for use data",
		"data": foruse,
	})
}