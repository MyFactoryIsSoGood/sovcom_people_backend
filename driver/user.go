package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
)

func GetUserById(id int) (bool, *models.User) {
	var user models.User

	db.DB.Model(&models.User{}).Where("id = ?", id).Scan(&user)
	if user.Email != "" {
		return true, &user
	}

	return false, &models.User{}
}
