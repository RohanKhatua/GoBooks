package routes

import (
	"context"
	"fmt"
	
	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/aws/aws-sdk-go-v2/aws"
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
	Item string `json:"item"`
	Path string `json:"path"`
	// DisplayProgress bool `json:"display_progress"`
}


