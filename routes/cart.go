package routes

import (
	"errors"
	"fmt"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RecvID struct {
	BookID uint `json:"book_id"`
}

func AddToCart(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var recvID RecvID

	if err := c.BodyParser(&recvID); err != nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	// check if book is already in cart

	var cartItem models.CartItem
	if err := database.Database.Db.Where("user_id = ? AND book_id = ?", userID, recvID.BookID).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myLogger.Info("No matching record found")
		} else {
			myLogger.Error("DB Search Failed: " + err.Error())
			return c.Status(400).JSON("DB Search Failed")
		}
	}

	if cartItem.ID != 0 {
		return c.Status(400).JSON("Book already in cart")
	}

	cartItem = models.CartItem{
		UserID: uint(userID),
		BookID: recvID.BookID,
	}

	err := database.Database.Db.Create(&cartItem).Error

	if err != nil {
		myLogger.Error("DB Insertion Failed")
		return c.Status(400).JSON(err.Error())
	}
	response := fmt.Sprintf("Added Book ID %d to Cart", recvID.BookID)

	return c.Status(200).JSON(response)

}

func GetCartItems(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var cartItems []models.CartItem
	if err := database.Database.Db.Where("user_id=?", userID).Find(&cartItems).Error; err != nil {
		if err.Error() == "record not found" {
			// myLogger.Info("No matching record found")
			return c.Status(400).JSON("No items in cart")
		}
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if len(cartItems) == 0 {
		// myLogger.Warning("Empty Cart")
		return c.Status(200).JSON("No items in cart")
	}
	return c.JSON(cartItems)

}

// Remove a book from cart by BookID
func RemoveFromCart(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))
	var recvID RecvID

	if err := c.BodyParser(&recvID); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var cartItem models.CartItem
	if err := database.Database.Db.Where("user_id = ? AND book_id = ?", userID, recvID.BookID).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myLogger.Info("No matching record found")
			return c.Status(400).JSON("Book not in cart")
		} else {
			myLogger.Error("DB Search Failed: " + err.Error())
			return c.Status(400).JSON("DB Search Failed")
		}
	}

	if err := database.Database.Db.Delete(&cartItem).Error; err != nil {
		myLogger.Error("DB Deletion Failed")
		return c.Status(400).JSON(err.Error())
	}

	response := fmt.Sprintf("Book ID %d removed", recvID.BookID)

	return c.Status(200).JSON(response)
}
