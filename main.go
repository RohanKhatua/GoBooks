package main

import (
	"log"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to GoAuth")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/signup", routes.SignUp)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
