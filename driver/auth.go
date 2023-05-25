package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
)

func CheckCredentials(phone, pass string) (bool, models.User) {
	var user models.User
	db.DB.Model(&models.User{}).Select("WHERE phone = ? AND password = ?", phone, pass).First(&user)

	if user.Phone != "" {
		return true, user
	} else {
		return false, models.User{}
	}
}
