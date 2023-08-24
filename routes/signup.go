package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/gofiber/fiber/v2"

	"github.com/RohanKhatua/fiber-jwt/helpers"
)

type UserReceive struct {
	UserName           string `json:"user_name"`
	Password           string `json:"pass"`
	SuperSecretAttempt string `json:"pk"`
}

type ResponseUser struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func CreateResponseUser(user models.User) ResponseUser {
	return ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
		Role:     user.Role,
	}
}

// get user info through request body
func SignUp(c *fiber.Ctx) error {
	var recdUserData UserReceive

	// problem with request body
	if err := c.BodyParser(&recdUserData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// New user based on User Model
	var newUser models.User

	if helpers.FindUserName(recdUserData.UserName) {
		log.Println("Duplicate User found")
		return c.Status(409).JSON("User Already Exists")
	}

	// Check if user has tried to get admin account and verify creds
	if recdUserData.SuperSecretAttempt != "" {
		super_secret := helpers.GetSuperSecret()
		fmt.Println("Reached Here")
		if super_secret == recdUserData.SuperSecretAttempt {
			log.Println("Key Matched, Admin Account Created")
			newUser.Role = "ADMIN"
		} else {
			log.Println("Bad Key - User Creation Failed")
			return c.Status(401).JSON("Unauthorized")
		}
	} else {
		newUser.Role = "USER"
	}

	newUser.UserName = recdUserData.UserName

	// Hash the password using SHA256
	hash := sha256.New()
	hash.Write([]byte(recdUserData.Password))
	hashBytes := hash.Sum(nil)
	newUser.Password = hex.EncodeToString(hashBytes)

	// Put the new user in the db

	database.Database.Db.Create(&newUser)

	// create JWT

	token, err := helpers.GenerateJWT(newUser)
	if err != nil {
		return c.Status(500).JSON("Internal Server Error")
	}

	// Response to return
	newResponseUser := CreateResponseUser(newUser)
	newResponseUser.Token = token

	return c.Status(200).JSON(newResponseUser)
}
