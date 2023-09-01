package routes

import (
	"log"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func ActivationMiddleware(c *fiber.Ctx) error {
	log.Println("From Middleware" , c.Locals("isActivated").(string))

	var user models.User

	var userID int = int(c.Locals("user_id").(float64))

	err:= database.Database.Db.Find(&user, "id=?", userID).Error

	if err != nil {
		log.Println("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if !user.IsActivated {
		log.Println("User not activated")
		return c.Status(400).JSON("User not activated")
	}

		return c.Next()
}