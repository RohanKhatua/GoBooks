package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type LoginData struct {
	UserName string `json:"user_name"`
	Password string `json:"pass"`
}

func Login(c *fiber.Ctx) error {
	myLogger := customLogger.NewLogger()
	var recdUserData LoginData

	if err:= c.BodyParser(&recdUserData); err!=nil {
		myLogger.Error("Failed to parse JSON")
		return c.Status(400).JSON(err.Error())
	}

	user,err := helpers.FindUserByName(recdUserData.UserName)

	if err!=nil {
		myLogger.Warning("Invalid Credentials - User does not exist")
		return c.Status(401).JSON("Invalid Credentials")
	}

	hash := sha256.New()
	hash.Write([]byte(recdUserData.Password))
	hashBytes := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashBytes)

	if user.Password != hashedPassword {
		myLogger.Warning("Invalid Credentials - UserName/Password is incorrect")
		return c.Status(401).JSON("Invalid Credentials")
	}

	log.Info("User Logged In, UserName: " + user.UserName)

	token, err := helpers.GenerateJWT(user)

	if err!=nil {
		myLogger.Error("Token Generation Failed")
		return c.Status(500).JSON("Internal Server Error")
	}

	newResponseUser := CreateResponseUser(user)
	newResponseUser.Token = token

	return c.Status(200).JSON(newResponseUser)
}