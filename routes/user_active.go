package routes

import (
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func ActivateUser(c *fiber.Ctx) error {
	var userID int = int(c.Locals("user_id").(float64))

	var user models.User

	database.Database.Db.Find(&user, "id=?", userID)

	if user.ID == 0 {
		return c.Status(400).JSON("Invalid User ID")
	}

	user.IsActivated = true

	database.Database.Db.Save(&user)	

	return c.Status(200).JSON("Account Activated")
}

func DeactivateUser(c *fiber.Ctx) error {
	var userID int = int(c.Locals("user_id").(float64))

	var user models.User

	database.Database.Db.Find(&user, "id=?", userID)

	if user.ID == 0 {
		return c.Status(400).JSON("Invalid User ID")
	}

	user.IsActivated = false

	database.Database.Db.Save(&user)	

	return c.Status(200).JSON("Account Deactivated")
}