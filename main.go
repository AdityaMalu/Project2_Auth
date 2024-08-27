package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/AdityaMalu/auth/data"
)

func main() { 
    data.Connect()

    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    app.Listen(":3000")
}