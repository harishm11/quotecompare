package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/urls"
)

func main() {
	app := fiber.New()
	urls.SetupRoutes(app)
	database.InitDatabase()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Listen(":8000")

}
