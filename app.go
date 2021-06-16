package main

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/train/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	bookRouter := app.Group("/api/v1")
	routes.ConfigBookRouter(&bookRouter)

	reviewRouter := app.Group("/api/v1")
	routes.ConfigReviewRouter(&reviewRouter)

	app.Listen(":3000")
}
