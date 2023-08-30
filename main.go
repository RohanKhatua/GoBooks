package main

import (
	"log"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("Welcome to the API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	// user routes
	app.Post("/api/signup", routes.SignUp)
	app.Post("/api/login", routes.Login)

	// implement middleware
	app.Use(routes.JWTMiddleware)
	app.Get("/api/protected", routes.ExampleProtectedRoute)

	//user activation routes

	app.Put("/api/activate", routes.ActivateUser)
	app.Put("/api/deactivate", routes.DeactivateUser)

	//book routes
	app.Post("/api/book", routes.CreateBook) //admin only route
	app.Get("/api/book/:id", routes.GetBookDetails)
	app.Get("/api/book", routes.GetBooks)
	app.Delete("/api/book", routes.DeleteBook) //admin only route
	app.Put("/api/book", routes.UpdateBook) //admin only route
	app.Put("/api/book/quantity", routes.ChangeBookQuantity) //admin only route

	// cart routes

	app.Post("/api/cart", routes.AddToCart)
	app.Get("/api/cart", routes.GetCartItems)
	app.Delete("/api/cart", routes.RemoveFromCart)

	// Review routes

	app.Post("/api/review", routes.AddReview)
	app.Get("/api/review/:id", routes.GetBookReviews) //reviews for a book
	app.Get("/api/review", routes.GetUserReviews) //reviews made by the user
	app.Put("/api/review", routes.EditReview)
	app.Delete("/api/review", routes.DeleteReview)
	app.Get("/api/rating/:id", routes.GetAverageRating)

	// Purchase routes

	app.Post("/api/purchase", routes.MakePurchase)
	app.Get("/api/purchase", routes.GetPurchases)

	// S3 routes

	app.Post("/api/upload", routes.UploadFile)
	app.Get("/api/download", routes.DownloadFile)

}

func main() {
	database.ConnectDb()
	// database.SeedDatabase(database.Database.Db)
	// database.CleanDatabase(database.Database.Db)
	app := fiber.New()

	setupRoutes(app)
	// log.Println("Back to Main")

	log.Fatal(app.Listen(":3000"))
}

