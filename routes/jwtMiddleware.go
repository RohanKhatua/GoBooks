package routes

import (
	"log"
	"time"

	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON("No Token")
	}

	//Remove the bearer prefix
	token = token[7:]

	// Validate the token
	secretKey := []byte(helpers.GetSuperSecret())
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !parsedToken.Valid || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return c.Status(401).JSON("Token Invalid")
	}

	// Set user ID in the context's "locals"
	c.Locals("user_id", claims["user_id"])
	c.Locals("user_name", claims["user_name"])
	c.Locals("user_role", claims["role"])
	log.Println("Current Logged In User :")
	log.Println("USER ID :",claims["user_id"])
	log.Println("USER NAME : ", claims["user_name"])
	log.Println("USER_ROLE : ", claims["role"])
	// Convert to int

	return c.Next()
}
