package routes

import (
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type RecvReview struct {
	BookID uint `json:"book_id"`
	Rating int `json:"rating"`
	Comment string `json:"comment"`
}

type RespReview struct {
	Rating int `json:"rating"`
	Comment string `json:"comment"`
}

func AddReview(c *fiber.Ctx) error {
	var userID int = int(c.Locals("user_id").(float64))

	var recvReview RecvReview

	if err:= c.BodyParser(&recvReview); err!=nil {
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", recvReview.BookID)

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	if recvReview.Rating < 1 || recvReview.Rating > 5 {
		return c.Status(400).JSON("Invalid Review : Rating must be between 1 and 5")
	} 

	var review models.Review = models.Review{
		UserID: uint(userID),
		BookID: recvReview.BookID,
		Rating: recvReview.Rating,
		Comment: recvReview.Comment,
	}

	database.Database.Db.Create(&review)

	respReview := RespReview {
		Rating: review.Rating,
		Comment: recvReview.Comment,
	}

	return c.Status(200).JSON(respReview)
}

func GetBookReviews (c* fiber.Ctx) error {
	// var userID int = int(c.Locals("user_id").(float64))
	
	bookID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("ID must be an integer")
	}

	var book models.Book

	database.Database.Db.Find(&book, "id=?", bookID)

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	var reviews []models.Review

	database.Database.Db.Find(&reviews, "book_id=?", bookID)

	var respReviews []RespReview

	for _, review := range reviews {
		respReview := RespReview {
			Rating: review.Rating,
			Comment: review.Comment,
		}
		respReviews = append(respReviews, respReview)
	}

	if len(respReviews) == 0 {
		return c.Status(400).JSON("No reviews found")
	}

	return c.Status(200).JSON(respReviews)
}

var GetUserReviews = func(c *fiber.Ctx) error {
	var userID int = int(c.Locals("user_id").(float64))

	var reviews []models.Review

	database.Database.Db.Find(&reviews, "user_id=?", userID)

	var respReviews []RespReview

	for _, review := range reviews {
		respReview := RespReview {
			Rating: review.Rating,
			Comment: review.Comment,
		}
		respReviews = append(respReviews, respReview)
	}

	if len(respReviews) == 0 {
		return c.Status(400).JSON("No reviews found")
	}

	return c.Status(200).JSON(respReviews)
}