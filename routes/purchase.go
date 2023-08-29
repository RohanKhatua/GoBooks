package routes

import (
	"fmt"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type RecvPurchase struct {
	BookID   uint `json:"book_id"`
	Quantity uint `json:"quantity"`
}

type ResponsePurchase struct {
	PurchaseID uint `json:"purchase_id"`
	BookID     uint `json:"book_id"`
	Quantity   uint `json:"quantity"`
	Price      uint `json:"price"`
}

func CreateResponsePurchase(purchase models.Purchase) ResponsePurchase {
	// get the price by multiplying the quantity with the price of the book
	myLogger := customLogger.NewLogger()
	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", purchase.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return ResponsePurchase{}
	}

	if book.ID == 0 {
		return ResponsePurchase{}
	}

	var price uint = book.Price * purchase.Quantity

	return ResponsePurchase{
		PurchaseID: purchase.ID,
		BookID:     purchase.BookID,
		Quantity:   purchase.Quantity,
		Price:      price,
	}
}

func MakePurchase(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var recvPurchase RecvPurchase

	if err := c.BodyParser(&recvPurchase); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", recvPurchase.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	if book.Quantity < recvPurchase.Quantity {
		resp := fmt.Sprintf("Not enough books in stock. Only %d books available", book.Quantity)

		return c.Status(400).JSON(resp)
	}

	var purchase models.Purchase = models.Purchase{
		UserID:   userID,
		BookID:   recvPurchase.BookID,
		Quantity: recvPurchase.Quantity,
	}

	err = database.Database.Db.Create(&purchase).Error

	if err != nil {
		myLogger.Error("DB Insertion Failed")
		return c.Status(400).JSON(err.Error())
	}

	book.Quantity -= recvPurchase.Quantity

	err = database.Database.Db.Save(&book).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	responsePurchase := CreateResponsePurchase(purchase)

	return c.Status(200).JSON(responsePurchase)

}

func GetPurchases(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var purchases []models.Purchase

	err := database.Database.Db.Find(&purchases, "user_id=?", userID).Error

	if err != nil {
		if err.Error() == "record not found" {
			// No purchases Made
			return c.Status(400).JSON("No purchases found")
		} else {
			myLogger.Error("DB Search Failed")
			return c.Status(400).JSON(err.Error())
		}
	}

	var responsePurchases []ResponsePurchase

	for _, purchase := range purchases {
		responsePurchases = append(responsePurchases, CreateResponsePurchase(purchase))
	}

	return c.Status(200).JSON(responsePurchases)
}
