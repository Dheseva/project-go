package controller

import (
	"project-go/database"
	"project-go/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterGun(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}

	airsoftgundata := models.AirsoftgunData{
		Typeammo: data["typeammo"],
		Capacity: data["capacity"],
		Material: data["material"],
		Color: data["color"],
		Seri_no: 01,
		Made_at: time.Now().UnixMilli(),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	database.DB.Create(&airsoftgundata)

	airsoftgun := models.Airsoftgun{
		Name: data["name"],
		Type: data["type"],
		From: data["from"],
		ASGData_id: int(airsoftgundata.Id),
		Available: 1,
		Created_by: 2,
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	database.DB.Create(&airsoftgun)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success register airsoftgun data",
		"data": fiber.Map{
			"airsoftgun": airsoftgun,
			"airsoftgun_data": airsoftgundata,
		},
	})
	
}