package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/AdityaMalu/Project2_Login_Auth/data"
)

func main() { 
	data.connect()

    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    app.Listen(":3000")
}