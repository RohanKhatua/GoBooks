package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var recdUserData UserReceive

	if err:= c.BodyParser(&recdUserData); err!=nil {
		return c.Status(400).JSON(err.Error())
	}

	user,err := helpers.FindUserByName(recdUserData.UserName)

	if err!=nil {
		return c.Status(401).JSON("Invalid Credentials")
	}

	hash := sha256.New()
	hash.Write([]byte(recdUserData.Password))
	hashBytes := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashBytes)

	if user.Password != hashedPassword {
		return c.Status(401).JSON("Invalid Credentials")
	}

	token, err := helpers.GenerateJWT(user)

	if err!=nil {
		return c.Status(500).JSON("Internal Server Error")
	}

	newResponseUser := CreateResponseUser(user)
	newResponseUser.Token = token

	return c.Status(200).JSON(newResponseUser)
}