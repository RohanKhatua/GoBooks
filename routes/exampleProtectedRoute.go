package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ExampleProtectedRoute(c *fiber.Ctx) error {
	log.Println(c.Locals("user_id").(float64))	

	var idInt int = int(c.Locals("user_id").(float64))
	var userRole string = c.Locals("user_role").(string)

	response := fmt.Sprintf("ID=%d Role=%s", idInt, userRole)
	
	return c.Status(200).JSON(response)
}