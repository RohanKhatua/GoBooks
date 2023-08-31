package routes

import (
	"errors"
	"log"
	"strings"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func searchBooks(searchTerm string) ([]models.Book, error) {
	mylogger := customLogger.NewLogger()
	books:= []models.Book{}

	searchTerm = strings.ToLower(searchTerm)

	log.Println("Search Term: " + searchTerm)

	if err := database.Database.Db.
		Where("LOWER(title) LIKE ? OR LOWER(author) LIKE ?", "%"+searchTerm+"%").
		Find(&books).
		Error; err != nil {
		return books, err
	}

	if len(books) == 0 {
		mylogger.Info("No matching records found")
		return books, errors.New("no matching records found")
	}

	return books, nil
}

func Search(c *fiber.Ctx) error {
	searchTerm := c.Query("q") // Get search query from query parameter

	log.Println("Search Term: ", searchTerm)

	books, err := searchBooks(searchTerm)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(books)
}
