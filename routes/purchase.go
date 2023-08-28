package routes

import (
	"fmt"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type RecvPurchase struct {
	BookID uint `json:"book_id"`
	Quantity uint `json:"quantity"`	
}

func MakePurchase(c *fiber.Ctx) error {
	var userID int = int(c.Locals("user_id").(float64))

	var recvPurchase RecvPurchase

	if err := c.BodyParser(&recvPurchase); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", recvPurchase.BookID)

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	if book.Quantity < recvPurchase.Quantity {
		resp := fmt.Sprintf("Not enough books in stock. Only %d books available", book.Quantity)

		return c.Status(400).JSON(resp)
	}

	var purchase models.Purchase = models.Purchase{
		UserID: userID,
		BookID: recvPurchase.BookID,
		Quantity: recvPurchase.Quantity,
	}

	database.Database.Db.Create(&purchase)

	book.Quantity -= recvPurchase.Quantity

	database.Database.Db.Save(&book)

	return c.Status(200).JSON("Purchase Successful")


}