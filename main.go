package main

import (
	"log"
	"os"

	"github.com/RianIhsan/ApiGoJwt/database"
	"github.com/RianIhsan/ApiGoJwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.InitRoute(app)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
