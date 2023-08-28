package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/RohanKhatua/fiber-jwt/customLogger"
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
	myLogger := customLogger.NewLogger()
	var recdUserData UserReceive

	// problem with request body
	if err := c.BodyParser(&recdUserData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// New user based on User Model
	var newUser models.User

	if helpers.FindUserName(recdUserData.UserName) {
		// log.Println("Duplicate User found")
		myLogger.Error("User Already Exists")
		return c.Status(409).JSON("User Already Exists")
	}

	// Check if user has tried to get admin account and verify creds
	if recdUserData.SuperSecretAttempt != "" {
		super_secret := helpers.GetSuperSecret()
		// fmt.Println("Reached Here")
		hash:= sha256.New()
		hash.Write([]byte(recdUserData.SuperSecretAttempt))
		hashed_attempt := hex.EncodeToString(hash.Sum(nil))

		if super_secret == hashed_attempt {
			myLogger.Info("Key Matched, Admin Account Created")
			newUser.Role = "ADMIN"
		} else {
			myLogger.Error("Bad Key - Admin Account Creation Failed")
			return c.Status(401).JSON("Unauthorized")
		}
	} else {
		newUser.Role = "USER"
	}

	newUser.UserName = recdUserData.UserName
	newUser.IsActivated = true

	// Hash the password using SHA256
	hash := sha256.New()
	hash.Write([]byte(recdUserData.Password))
	hashBytes := hash.Sum(nil)
	newUser.Password = hex.EncodeToString(hashBytes)

	// Put the new user in the db

	database.Database.Db.Create(&newUser)
	myLogger.Info("Added New User to DB")

	// create JWT

	token, err := helpers.GenerateJWT(newUser)
	if err != nil {
		myLogger.Error("Token Creation Failed")
		return c.Status(500).JSON("Internal Server Error")
	}

	// Response to return
	newResponseUser := CreateResponseUser(newUser)
	newResponseUser.Token = token

	return c.Status(200).JSON(newResponseUser)
}
