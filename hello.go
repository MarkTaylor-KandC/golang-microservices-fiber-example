package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// start the application on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	return app
}
