package routes

import "github.com/gofiber/fiber/v2"

func ExampleProtectedRoute(c *fiber.Ctx) error {
	return c.Status(200).JSON("This is a protected Route")
}