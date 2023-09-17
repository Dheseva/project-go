package main

import (
	"os"
	"project-go/database"
	"project-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	
	// loadEnv()
	database.Connect()
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(getPort())
}

func getPort() string{
	port := os.Getenv("DB_PORT")
	if port == ""{
		return ":8000"
	}
	return ":" + port
}

// func loadEnv(){
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	  }
	 
// }