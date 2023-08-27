package helpers

func IsAdmin(user_id uint) (bool, error) {
	user, err := FindUserByID(uint(user_id))

	if err!=nil {
		return false, err
	}

	return user.Role == "ADMIN", nil
}