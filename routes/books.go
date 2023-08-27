package routes

import (
	"log"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/helpers"
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
	userRole := c.Locals("user_role").(string)
	isAdmin := userRole == "ADMIN"

	log.Println("ROLE : ", userRole)

	if !isAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Must be an admin",
		})
	}

	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
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
	curr_user_id := c.Locals("user_id").(uint)

	isAdmin, err := helpers.IsAdmin(curr_user_id)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if !isAdmin {
		return c.Status(401).JSON("Must be Admin")
	}

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("ID must be an integer")
	}

	if id == 0 {
		return c.Status(400).JSON("Book does not exist")
	}

	var book models.Book

	database.Database.Db.Delete(&book, "id=?", id)

	return c.Status(200).JSON("Book Deleted")
}
