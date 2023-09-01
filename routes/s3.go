package routes

import (
	"context"
	"fmt"

	"os"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var user_role = c.Locals("user_role")

	if user_role != "ADMIN" {
		myLogger.Warning("Unauthorized access to /api/upload")
		return c.Status(401).JSON("Unauthorized")
	}

	// File Name of the form <Author> - <Title>.pdf
	file, err := c.FormFile("uploadFile")

	if err != nil {
		return c.Status(400).JSON("Error fetching file")
	}

	// Open File

	uploadFile, err := file.Open()

	if err != nil {
		return c.Status(400).JSON("Error opening file")
	}

	uploader := helpers.CreateUploader()
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("balkanid-api-book-storage"),
		Key:    aws.String(file.Filename),
		Body:   uploadFile,
		ACL:    "public-read",
	})

	if err != nil {
		return c.Status(400).JSON("Error uploading file")
	}

	fmt.Println(result)

	return c.Status(200).JSON("File uploaded successfully")
}

type DownloadRequest struct {
	Item string `json:"item"` //filename
	Path string `json:"path"` //path to download to
}

func createDownloadRequest (book_id uint) (DownloadRequest, error) {
	var book models.Book

	if err:= database.Database.Db.Where("id = ?",book_id).First(&book).Error; err!=nil {
		return DownloadRequest{}, err
	}

	author := book.Author
	title := book.Title

	filename := fmt.Sprintf("%s - %s.pdf",author,title)

	return DownloadRequest{
		Item: filename,
		Path: "./",
	}, nil
}

func DownloadFile(c *fiber.Ctx) error {
	var recvID RecvID

	if err:= c.BodyParser(&recvID); err!=nil {
		return c.Status(400).JSON("Failed to Parse JSON")
	}	

	// Check if user has purchased book

	var userID int = int(c.Locals("user_id").(float64))

	var purchase models.Purchase

	if err:= database.Database.Db.Where("user_id = ? AND book_id = ?",userID,recvID.BookID).First(&purchase).Error; err!=nil {
		return c.Status(400).JSON("Book not purchased")
	}

	// Create Download Request

	downloadRequest, err := createDownloadRequest(recvID.BookID)

	if err!=nil {
		return c.Status(400).JSON("Failed to Create Download Request")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err!=nil {
		return c.Status(500).JSON("Failed to Create Configuration")
	}

	client := s3.NewFromConfig(cfg)

	downloader := manager.NewDownloader(client)
	

	file, err := os.Create(downloadRequest.Item)
	if err!=nil {
		return c.Status(400).JSON("Failed to Create File")
	}

	defer file.Close()

	_, err = downloader.Download(context.TODO(),file,&s3.GetObjectInput{
		Bucket: aws.String("balkanid-api-book-storage"),
		Key: &downloadRequest.Item,		
	})

	if err!=nil {
		return c.Status(400).JSON("Failed to download file")
	}

	defer os.Remove(downloadRequest.Item)

	return c.Status(200).SendFile(downloadRequest.Item)

	// return c.Status(200).JSON("File Downloaded" + downloadRequest.Path)
}
