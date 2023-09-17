package main

import (
	"log"
	"project-go/database"
	"project-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	
	loadEnv()
	database.Connect()
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}

func loadEnv(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	  }
	 
}