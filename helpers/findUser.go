package helpers

import (
	"errors"

	"github.com/RohanKhatua/fiber-jwt/database"
	"github.com/RohanKhatua/fiber-jwt/models"
)

//this is a helper function not a route

//find by userName check if the username is already taken
func FindUserName (userName string) bool{
	var user models.User
	database.Database.Db.First(&user, "user_name=?", userName)
	return user.ID!=0 
}

func FindUserByName (userName string) (models.User, error) {
	var user models.User
	database.Database.Db.First(&user, "user_name=?", userName)

	if user.ID==0 {
		return user, errors.New("user does not exist")
	}

	return user, nil
}