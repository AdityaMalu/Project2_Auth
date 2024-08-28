package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/AdityaMalu/auth/database"
    "github.com/AdityaMalu/auth/routes"
)

func main() { 
    database.Connect()

    app := fiber.New()

    routes.Setup(app)

    app.Listen(":3000")
}