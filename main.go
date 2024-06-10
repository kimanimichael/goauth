package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mike-kimani/goauth/database"
	"github.com/mike-kimani/goauth/routes"
	"log"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowCredentials: true,
	}))

	routes.Setup(app)
	routes.Register(app)
	routes.Login(app)

	err := app.Listen(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
