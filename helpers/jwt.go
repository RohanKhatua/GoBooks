package helpers

import (
	"time"

	"github.com/RohanKhatua/fiber-jwt/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//Set claims
	claims:= token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["user_name"] = user.UserName
	claims["role"] = user.Role

	expirationTime := time.Now().Add(24*time.Hour)
	claims["exp"] = expirationTime.Unix()

	secretKey := []byte(GetSuperSecret())
	tokenString, err := token.SignedString(secretKey)
	if err!=nil {
		return "", err
	}

	return tokenString, nil
}
