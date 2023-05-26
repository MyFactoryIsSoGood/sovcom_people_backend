package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"errors"
)

func CheckCredentials(email, pass string) (bool, *models.User) {
	var user models.User
	db.DB.Model(&models.User{}).Where("email = ? AND password = ?", email, pass).First(&user)

	if user.Email != "" {
		return true, &user
	}
	return false, &models.User{}
}

func GetUserByMail(email string) (bool, *models.User) {
	var user models.User

	db.DB.Model(&models.User{}).Where("email = ?", email).Scan(&user)

	if user.Email != "" {
		return true, &user
	}
	return false, &models.User{}
}

func CreateUser(phone, email, password, fullName string, role uint) (error, *models.User) {
	if found, user := GetUserByMail(email); found {
		return errors.New("already exists"), user
	}

	var user = models.User{
		FullName: fullName,
		Role:     role,
		Phone:    phone,
		Email:    email,
		Password: password,
	}
	err := db.DB.Model(&models.User{}).Create(&user).Error

	if err != nil {
		return err, nil
	}
	return nil, &user
}
