package routes

import (
	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func ActivateUser(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var user models.User

	err:= database.Database.Db.Find(&user, "id=?", userID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if user.ID == 0 {
		return c.Status(400).JSON("Invalid User ID")
	}

	user.IsActivated = true
	// c.Locals("isActivated", "true")

	err = database.Database.Db.Save(&user).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Account Activated")
}

func DeactivateUser(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var user models.User

	err:= database.Database.Db.Find(&user, "id=?", userID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if user.ID == 0 {
		return c.Status(400).JSON("Invalid User ID")
	}

	user.IsActivated = false
	// c.Locals("isActivated", "false")

	err = database.Database.Db.Save(&user).Error
	
	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Account Deactivated")
}