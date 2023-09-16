package controller

import (
	"project-go/database"
	"project-go/models"
	"strconv"
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

func AirsoftgunView(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	var airsoftgun models.Airsoftgun
	var airsoftgundata models.AirsoftgunData

	database.DB.Where("id = ?", idInt).First(&airsoftgun)
	database.DB.Where("id = ?", airsoftgun.ASGData_id).First(&airsoftgundata)

	if airsoftgun.Id == 0 || airsoftgundata.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "data not found",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success get data",
		"data": fiber.Map{
			"airsoftgun" : airsoftgun,
			"airsoftgun_data": airsoftgundata,
		},
	})
}

func AirsoftgunUpdate(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}
	
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	var airsoftgun models.Airsoftgun
	var airsoftgundata models.AirsoftgunData

	database.DB.Where("id = ?", idInt).First(&airsoftgun)
	database.DB.Where("id = ?", airsoftgun.ASGData_id).First(&airsoftgundata)

	if airsoftgun.Id == 0 || airsoftgundata.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "data not found",
			"data": nil,
		})
	}

	airsoftgun = models.Airsoftgun{
		Id: airsoftgun.Id,
		Name: data["name"],
		Type: data["type"],
		From: data["from"],
		ASGData_id: airsoftgun.ASGData_id,
		Available: airsoftgun.Available,
		Created_by: airsoftgun.Created_by,
		Created_at: airsoftgun.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: airsoftgun.Deleted_at,
	}

	airsoftgundata = models.AirsoftgunData{
		Id: airsoftgundata.Id,
		Typeammo: data["typeammo"],
		Capacity: data["capacity"],
		Material: data["material"],
		Color: data["color"],
		Seri_no: airsoftgundata.Seri_no,
		Made_at: airsoftgundata.Made_at,
		Created_at: airsoftgundata.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: airsoftgundata.Deleted_at,
	}

	database.DB.Save(&airsoftgun)
	database.DB.Save(&airsoftgundata)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success update data",
		"data": fiber.Map{
			"airsoftgun" : airsoftgun,
			"airsoftgun_data": airsoftgundata,
		},
	})
}

func AirsoftgunDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	var airsoftgun models.Airsoftgun
	var airsoftgundata models.AirsoftgunData

	database.DB.Where("id = ?", idInt).First(&airsoftgun)
	database.DB.Where("id = ?", airsoftgun.ASGData_id).First(&airsoftgundata)

	if airsoftgun.Id == 0 || airsoftgundata.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "data not found",
			"data": nil,
		})
	}

	airsoftgun = models.Airsoftgun{
		Id: airsoftgun.Id,
		Name: airsoftgun.Name,
		Type: airsoftgun.Type,
		From: airsoftgun.From,
		ASGData_id: airsoftgun.ASGData_id,
		Available: airsoftgun.Available,
		Created_by: airsoftgun.Created_by,
		Created_at: airsoftgun.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}

	airsoftgundata = models.AirsoftgunData{
		Id: airsoftgundata.Id,
		Typeammo: airsoftgundata.Typeammo,
		Capacity: airsoftgundata.Capacity,
		Material: airsoftgundata.Material,
		Color: airsoftgundata.Color,
		Seri_no: airsoftgundata.Seri_no,
		Made_at: airsoftgundata.Made_at,
		Created_at: airsoftgundata.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}

	database.DB.Save(&airsoftgun)
	database.DB.Save(&airsoftgundata)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success delete data",
		"data": fiber.Map{
			"airsoftgun" : airsoftgun,
			"airsoftgun_data": airsoftgundata,
		},
	})
}