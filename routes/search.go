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

func searchBooks(searchTerm string) ([]ResponseBook, error) {
	mylogger := customLogger.NewLogger()
	books:= []models.Book{}
	responseBooks := []ResponseBook{}

	searchTerm = strings.ToLower(searchTerm)

	log.Println("Search Term: " + searchTerm)

	if err := database.Database.Db.
		Where("LOWER(title) LIKE ? OR LOWER(author) LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").
		Find(&books).
		Error; err != nil {
		return responseBooks, err
	}

	if len(books) == 0 {
		mylogger.Info("No matching records found")
		return responseBooks, errors.New("no matching records found")
	}


	for _, book := range books {
		responseBook := CreateResponseBook(book)
		responseBooks = append(responseBooks, responseBook)
	}

	return responseBooks, nil
}

type SearchRequest struct {
	SearchTerm string `json:"query"`
}

func Search(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()

	var searchRequest SearchRequest

	err := c.BodyParser(&searchRequest)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	searchTerm := searchRequest.SearchTerm

	// All spaces are replaced with %20 in the URL

	// searchTerm = strings.ReplaceAll(searchTerm, " ", "%20")

	myLogger.Info("Search Term: " + searchTerm)

	books, err := searchBooks(searchTerm)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(books)
}
