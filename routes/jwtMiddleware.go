package routes

import (
	"time"

	"github.com/RohanKhatua/fiber-jwt/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON("Unauthorized")
	}

	// Validate the token
	secretKey := []byte(helpers.GetSuperSecret())
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !parsedToken.Valid || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return c.Status(401).JSON("Unauthorized")
	}

	// Set user ID in the context's "locals"
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(401).JSON("Unauthorized")
	}
	c.Locals("user_id", int(userID)) // Convert to int

	return c.Next()
}
