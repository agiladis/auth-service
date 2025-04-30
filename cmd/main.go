package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/agiladis/auth-service/config"
	"github.com/agiladis/auth-service/internal/api/v1/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	router.SetupWebRoutes(v1, db)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
