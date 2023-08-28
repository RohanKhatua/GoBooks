package routes

import (
	"fmt"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type ResponseBook struct {
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func CreateResponseBook(book models.Book) ResponseBook {

	return ResponseBook{
		Author:      book.Author,
		Year:        book.Year,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		//user does not need to know quantity
	}
}

func CreateBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to add book, UserName : "+c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	err = database.Database.Db.Create(&book).Error

	if err != nil {
		myLogger.Error("Adding Book to DB Failed")
		return c.Status(400).JSON(err.Error())
	}

	myLogger.Info("Book Added to DB, BookID:"+fmt.Sprint(book.ID))

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
		// myLogger.Warning("Book does not exist")
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	newResponseBook := CreateResponseBook(book)

	return c.Status(200).JSON(newResponseBook)
}

func GetBooks(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	books := []models.Book{}
	err:= database.Database.Db.Find(&books).Error

	if err!=nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}
	ResponseBooks := []ResponseBook{}

	for _, book := range books {
		responseBook := CreateResponseBook(book)
		ResponseBooks = append(ResponseBooks, responseBook)
	}

	if len(ResponseBooks) == 0 {
		myLogger.Warning("Empty DB")
		return c.Status(400).JSON("No Books Found")
	}

	return c.Status(200).JSON(ResponseBooks)
}

func DeleteBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to delete book, UserName : "+c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvID RecvID

	if err:=c.BodyParser(&recvID); err!=nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	if recvID.BookID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var book models.Book

	err := database.Database.Db.Delete(&book, "id=?", recvID.BookID).Error

	if err != nil {
		myLogger.Error("DB Delete Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Book Deleted, BookID:"+fmt.Sprint(book.ID))
}

func UpdateBook(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to update book, UserName : "+c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvID RecvID

	if err:=c.BodyParser(&recvID); err!=nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	if recvID.BookID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var book models.Book

	err:= database.Database.Db.Find(&book, "id=?", recvID.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var newBook models.Book

	if err := c.BodyParser(&newBook); err != nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	err = database.Database.Db.Model(&book).Updates(newBook).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Book Updated, BookID:"+fmt.Sprint(book.ID))
}

type RecvBookQuantity struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

var ChangeBookQuantity = func(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		myLogger.Warning("Non Admin User tried to change book quantity, UserName : "+c.Locals("user_name").(string))
		return c.Status(401).JSON("Must be Admin")
	}

	var recvBook RecvBookQuantity

	if err := c.BodyParser(&recvBook); err != nil {
		myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	err:= database.Database.Db.Find(&book, "id=?", recvBook.ID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	err = database.Database.Db.Model(&book).Update("quantity", recvBook.Quantity).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Book Quantity Updated, BookID:"+fmt.Sprint(book.ID)+", New Quantity:"+fmt.Sprint(book.Quantity))
}
