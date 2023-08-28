package routes

import (
	// "fmt"

	"math"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

type RecvReview struct {
	BookID  uint   `json:"book_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type RespReview struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func AddReview(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var recvReview RecvReview

	if err := c.BodyParser(&recvReview); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", recvReview.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	// find if user has purchased the book

	var purchase models.Purchase

	err = database.Database.Db.Find(&purchase, "user_id=? AND book_id=?", userID, recvReview.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if purchase.ID == 0 {
		return c.Status(400).JSON("You have not purchased this book")
	}

	// check if review already exists

	var checkReview models.Review

	err = database.Database.Db.Find(&checkReview, "user_id=? AND book_id=?", userID, recvReview.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if checkReview.ID != 0 {
		return c.Status(400).JSON("You have already reviewed this book")
	}

	if recvReview.Rating < 1 || recvReview.Rating > 5 {
		return c.Status(400).JSON("Invalid Review : Rating must be between 1 and 5")
	}

	var review models.Review = models.Review{
		UserID:  uint(userID),
		BookID:  recvReview.BookID,
		Rating:  recvReview.Rating,
		Comment: recvReview.Comment,
	}

	err = database.Database.Db.Create(&review).Error

	if err != nil {
		myLogger.Error("DB Insertion Failed")
		return c.Status(400).JSON(err.Error())
	}

	respReview := RespReview{
		Rating:  review.Rating,
		Comment: recvReview.Comment,
	}

	return c.Status(200).JSON(respReview)
}

func GetBookReviews(c *fiber.Ctx) error {
	// var userID int = int(c.Locals("user_id").(float64))
	myLogger := customLogger.NewLogger()
	bookID, err := c.ParamsInt("id")

	if err != nil {
		myLogger.Warning("Non Integer ID passed")
		return c.Status(400).JSON("ID must be an integer")
	}

	var book models.Book

	err = database.Database.Db.Find(&book, "id=?", bookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	var reviews []models.Review

	err = database.Database.Db.Find(&reviews, "book_id=?", bookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	var respReviews []RespReview

	for _, review := range reviews {
		respReview := RespReview{
			Rating:  review.Rating,
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
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var reviews []models.Review

	err := database.Database.Db.Find(&reviews, "user_id=?", userID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	var respReviews []RespReview

	for _, review := range reviews {
		respReview := RespReview{
			Rating:  review.Rating,
			Comment: review.Comment,
		}
		respReviews = append(respReviews, respReview)
	}

	if len(respReviews) == 0 {
		return c.Status(400).JSON("No reviews found")
	}

	return c.Status(200).JSON(respReviews)
}

func EditReview(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var recvReview RecvReview

	if err := c.BodyParser(&recvReview); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", recvReview.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	if recvReview.Rating < 1 || recvReview.Rating > 5 {
		return c.Status(400).JSON("Invalid Review : Rating must be between 1 and 5")
	}

	var review models.Review

	err = database.Database.Db.Find(&review, "user_id=? AND book_id=?", userID, recvReview.BookID).Error

	if err != nil {

		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())

	}

	if review.ID == 0 {
		return c.Status(400).JSON("No review found")
	}

	review.Rating = recvReview.Rating
	review.Comment = recvReview.Comment

	err = database.Database.Db.Save(&review).Error

	if err != nil {
		myLogger.Error("DB Update Failed")
		return c.Status(400).JSON(err.Error())
	}

	respReview := RespReview{
		Rating:  review.Rating,
		Comment: recvReview.Comment,
	}

	return c.Status(200).JSON(respReview)
}

var DeleteReview = func(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var userID int = int(c.Locals("user_id").(float64))

	var recvID RecvID

	if err := c.BodyParser(&recvID); err != nil {
		//myLogger.Error("JSON Parsing Failed")
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	err := database.Database.Db.Find(&book, "id=?", recvID.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if book.ID == 0 {
		return c.Status(400).JSON("Invalid Book ID")
	}

	var review models.Review

	err = database.Database.Db.Find(&review, "user_id=? AND book_id=?", userID, recvID.BookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if review.ID == 0 {
		return c.Status(400).JSON("No review found")
	}

	err = database.Database.Db.Delete(&review).Error

	if err != nil {
		myLogger.Error("DB Deletion Failed")
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("Review deleted")
}

func GetAverageRating(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	bookID, err := c.ParamsInt("id")

	if err != nil {
		myLogger.Warning("Non Integer ID passed")
		return c.Status(400).JSON("ID must be an integer")
	}

	var reviews []models.Review

	err = database.Database.Db.Find(&reviews, "book_id=?", bookID).Error

	if err != nil {
		myLogger.Error("DB Search Failed")
		return c.Status(400).JSON(err.Error())
	}

	if len(reviews) == 0 {
		return c.Status(400).JSON("No reviews found for this book")
	}

	var sum int = 0

	for _, review := range reviews {
		sum += review.Rating
	}

	var avg float64 = float64(sum) / float64(len(reviews))

	//round avg to 2 decimal places

	avg = math.Round(avg*100) / 100

	// resp := fmt.Sprintf("Average Rating for Book ID %d is %f", bookID, avg)

	return c.Status(200).JSON(avg)
}
