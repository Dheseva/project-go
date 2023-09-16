package controller

import (
	"project-go/database"
	"project-go/middleware"
	"project-go/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func UpdateDataUser(c *fiber.Ctx) error {
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

	cookie := c.Cookies("jwt")
	standClaims := &middleware.MyCustomClaims{}
	convKey := []byte(SecretKey)
	token, err := jwt.ParseWithClaims(cookie, standClaims, func(t *jwt.Token) (interface{}, error) {
		return convKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  false,
			"error":   err.Error(),
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*middleware.MyCustomClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	var userdata models.UserData
	database.DB.Where("id = ?", user.UData_id).First(&userdata)

	useridInt := int(user.Id)
	if idInt != useridInt {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "unauthenticated user",
		})
	}

	userdata = models.UserData{
		Id: userdata.Id,
		Fullname: data["fullname"],
		Sex: data["sex"],
		Address: data["address"],
		Dateofbirth: 998233200000,
		Nationality: data["nationality"],
		Created_at: userdata.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: userdata.Deleted_at,

	}

	database.DB.Save(&userdata)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success update data user",
		"data": fiber.Map{
			"user":      user,
			"user_data": userdata,
		},
	})
}