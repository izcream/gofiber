package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
		})
	})
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Hello " + c.Params("name"),
		})
	})
	app.Listen(":" + port)
}
