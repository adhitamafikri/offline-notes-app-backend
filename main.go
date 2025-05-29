package main

import (
	"fmt"
	"offline-notes-app-backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	fmt.Println("Staring server with host: ", host, "and port: ", port)

	app.Listen(host + ":" + port)
}
