package main

import (
	"go-auth/database"
	"go-auth/routes"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
}

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
