package controller

import (
	"project-go/database"
	"project-go/models"
	"strconv"
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

func ForUseView(c *fiber.Ctx) error {
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

	var foruse models.ForUse

	database.DB.Where("id = ?", idInt).First(&foruse)

	if foruse.Id == 0{
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
		"data": foruse,
	})
}

func ForUseUpdate(c *fiber.Ctx) error {
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

	var foruse models.ForUse

	database.DB.Where("id = ?", idInt).First(&foruse)

	if foruse.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "data not found",
			"data": nil,
		})
	}

	foruse = models.ForUse{
		Id: foruse.Id,
		User_id: foruse.User_id,
		Airsoftgun_id: foruse.Airsoftgun_id,
		Use_for: data["use_for"],
		Comment: data["comment"],
		Expired_at: foruse.Expired_at,
		Created_at: foruse.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: foruse.Deleted_at,
	}

	database.DB.Save(&foruse)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success update data",
		"data": foruse,
	})
}

func ForUseDelete(c *fiber.Ctx) error {
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

	var foruse models.ForUse

	database.DB.Where("id = ?", idInt).First(&foruse)

	if foruse.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "data not found",
			"data": nil,
		})
	}

	foruse = models.ForUse{
		Id: foruse.Id,
		User_id: foruse.User_id,
		Airsoftgun_id: foruse.Airsoftgun_id,
		Use_for: foruse.Use_for,
		Comment: foruse.Comment,
		Expired_at: foruse.Expired_at,
		Created_at: foruse.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}

	database.DB.Save(&foruse)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success delete data",
		"data": foruse,
	})
}