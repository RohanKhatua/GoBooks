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

	// user routes
	app.Post("/api/signup", routes.SignUp)
	app.Post("/api/login", routes.Login)

	// implement middleware
	app.Use(routes.JWTMiddleware)
	app.Get("/api/protected", routes.ExampleProtectedRoute)

	//book routes
	app.Get("/api/books", routes.GetBooks)
	app.Get("/api/books/:id", routes.GetBookDetails)
	app.Post("/api/books", routes.CreateBook)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
