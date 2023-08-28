package routes

import (
	"log"

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
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		return c.Status(401).JSON("Must be Admin")
	}

	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&book)

	log.Println("Added to DB")

	responseBook := CreateResponseBook(book)

	return c.Status(fiber.StatusOK).JSON(responseBook)
}

func GetBookDetails(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("ID must be an integer")
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", id)

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	newResponseBook := CreateResponseBook(book)

	return c.Status(200).JSON(newResponseBook)
}

func GetBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	database.Database.Db.Find(&books)
	ResponseBooks := []ResponseBook{}

	for _, book := range books {
		responseBook := CreateResponseBook(book)
		ResponseBooks = append(ResponseBooks, responseBook)
	}

	return c.Status(200).JSON(ResponseBooks)
}

func DeleteBook(c *fiber.Ctx) error {
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		return c.Status(401).JSON("Must be Admin")
	}

	var recvID RecvID

	if err:=c.BodyParser(&recvID); err!=nil {
		return c.Status(400).JSON(err.Error())
	}

	if recvID.BookID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var book models.Book

	database.Database.Db.Delete(&book, "id=?", recvID.BookID)

	return c.Status(200).JSON("Book Deleted")
}

func UpdateBook(c *fiber.Ctx) error {
	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		return c.Status(401).JSON("Must be Admin")
	}

	var recvID RecvID

	if err:=c.BodyParser(&recvID); err!=nil {
		return c.Status(400).JSON(err.Error())
	}

	if recvID.BookID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", recvID.BookID)

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var newBook models.Book

	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Model(&book).Updates(newBook)

	return c.Status(200).JSON("Book Updated")
}

type RecvBookQuantity struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

var ChangeBookQuantity = func(c *fiber.Ctx) error {

	var userRole string = c.Locals("user_role").(string)

	if userRole != "ADMIN" {
		return c.Status(401).JSON("Must be Admin")
	}

	var recvBook RecvBookQuantity

	if err := c.BodyParser(&recvBook); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", recvBook.ID)

	if book.ID == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	database.Database.Db.Model(&book).Update("quantity", recvBook.Quantity)

	return c.Status(200).JSON("Book Quantity Updated")
}
