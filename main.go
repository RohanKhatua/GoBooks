package main

import (
	"log"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	// below this all routes are protected by JWT Middleware have to be signed in to access them
	app.Get("/api/protected", routes.ExampleProtectedRoute)

	//user activation routes

	// activation route is unprotected by Activation Middleware
	app.Put("/api/activate", routes.ActivateUser)

	app.Use(routes.ActivationMiddleware)

	// below this all routes are protected by Activation Middleware have to be activated to access them

	// Cannot deactivate unless activated first
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

	// Search routes

	app.Get("/api/search", routes.Search)

}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load .env file")
	}


	database.ConnectDb()

	// database.CleanDatabase(database.Database.Db)
	// database.SeedDatabase(database.Database.Db)
	// database.SeedDatabase(database.Database.Db)
	app := fiber.New()

	setupRoutes(app)
	// log.Println("Back to Main")

	log.Fatal(app.Listen(":3000"))
}

