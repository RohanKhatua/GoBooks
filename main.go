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
	app.Delete("/api/book/:id", routes.DeleteBook)

	// cart routes

	app.Get("/api/cart", routes.GetCartItems)
	app.Post("/api/cart", routes.AddToCart)
	app.Delete("/api/cart", routes.RemoveFromCart)

	// Review routes

	app.Post("/api/review", routes.AddReview)
	app.Get("/api/review/:id", routes.GetBookReviews) //reviews for a book
	app.Get("/api/review", routes.GetUserReviews) //reviews made by the user

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
