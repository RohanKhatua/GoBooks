package routes

import (
	"fmt"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type ResponseBook struct {
	ID          uint   `json:"id"`
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func CreateResponseBook(book models.Book) ResponseBook {

	return ResponseBook{
		ID:          book.ID,
		Author:      book.Author,
		Year:        book.Year,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		//user does not need to know quantity
	}
}

type RecvUpdatedBook struct {
	BookID      uint   `json:"book_id"`
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func CreateBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to add book, UserName : " + c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	// check if book already exists

	var existingBook models.Book

	err = database.Database.Db.Find(&existingBook, "title=?", book.Title).Error

	if err != nil {
		if err.Error() == "record not found" {
			// book does not exist
		} else {
			myLogger.Error("DB Search Failed")
			return c.Status(400).JSON(err.Error())
		}
	}

	if existingBook.ID != 0 {
		return c.Status(400).JSON("Book already exists")
	}

	err = database.Database.Db.Create(&book).Error

	if err != nil {
		myLogger.Error("Adding Book to DB Failed")
		return c.Status(400).JSON(err.Error())
	}

	myLogger.Info("Book Added to DB, BookID:" + fmt.Sprint(book.ID))

	responseBook := CreateResponseBook(book)

	return c.Status(fiber.StatusOK).JSON(responseBook)
}

func GetBookDetails(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	id, err := c.ParamsInt("id")

	if err != nil {
		myLogger.Warning("Non Integer ID entered")
		return c.Status(400).JSON("ID must be an integer")
	}

	var book models.Book

	err = database.Database.Db.Find(&book, "id=?", id).Error

	if err != nil {
		if err.Error() == "record not found" {
			// myLogger.Warning("Book does not exist")
			return c.Status(400).JSON("Book does not exist")
		}
		// myLogger.Warning("Book does not exist")
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	// if book.ID == 0 {
	// 	return c.Status(400).JSON("Book does not exist")
	// }

	newResponseBook := CreateResponseBook(book)

	return c.Status(200).JSON(newResponseBook)
}

func GetBooks(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	books := []models.Book{}
	err := database.Database.Db.Find(&books).Error

	if err != nil {
		if err.Error() == "record not found" {
			// myLogger.Warning("Empty DB")
			return c.Status(400).JSON("No Books Found")
		}
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}
	ResponseBooks := []ResponseBook{}

	for _, book := range books {
		responseBook := CreateResponseBook(book)
		ResponseBooks = append(ResponseBooks, responseBook)
	}

	return c.Status(200).JSON(ResponseBooks)
}

func DeleteBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to delete book, UserName : " + c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvID RecvID

	if err := c.BodyParser(&recvID); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	resp := "Book Deleted, BookID:" + fmt.Sprint(recvID.BookID)

	var book models.Book

	err := database.Database.Db.Delete(&book, "id=?", recvID.BookID)

	// if delete does not affect any rows then book does not exist
	
	if err.RowsAffected == 0 {
		return c.Status(400).JSON("Book does not exist")
	}
	myLogger.Info("Rows Affected Without Error")
	return c.Status(200).JSON(resp)
}

func UpdateBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to update book, UserName : " + c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvUpdatedBook RecvUpdatedBook

	if err := c.BodyParser(&recvUpdatedBook); err != nil {
		return c.Status(400).JSON("Failed to Parse JSON")
	}

	// check if bookID exists

	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", recvUpdatedBook.BookID).Error

	if err != nil {
		if err.Error() == "record not found" {
			// myLogger.Warning("Book does not exist")
			return c.Status(400).JSON("Book does not exist")
		}
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	// update book

	err = database.Database.Db.Model(&book).Updates(models.Book{
		Author:      recvUpdatedBook.Author,
		Year:        recvUpdatedBook.Year,
		Title:       recvUpdatedBook.Title,
		Description: recvUpdatedBook.Description,
		Price:       recvUpdatedBook.Price,
	}).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Book Updated, BookID:" + fmt.Sprint(book.ID))
}

type RecvBookQuantity struct {
	ID       uint `json:"book_id"`
	Quantity uint `json:"quantity"`
}

var ChangeBookQuantity = func(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to change book quantity, UserName : " + c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvBook RecvBookQuantity

	if err := c.BodyParser(&recvBook); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	if err := database.Database.Db.Find(&book, "id=?", recvBook.ID).Error; err != nil {
		if err.Error() == "record not found" {
			// myLogger.Warning("Book does not exist")
			return c.Status(400).JSON("Book does not exist")
		}
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	err := database.Database.Db.Model(&book).Update("quantity", recvBook.Quantity).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Book Quantity Updated, BookID:" + fmt.Sprint(book.ID) + ", New Quantity:" + fmt.Sprint(book.Quantity))
}
