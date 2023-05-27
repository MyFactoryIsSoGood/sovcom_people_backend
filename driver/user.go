package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
)

func GetUserById(id int) (bool, *models.User) {
	var user models.User

	db.DB.Model(&models.User{}).Preload("CVs.Applies").Find(&user, id)
	if user.Email != "" {
		return true, &user
	}

	return false, &models.User{}
}

func GetUsers() []models.User {
	var users []models.User

	db.DB.Model(&models.User{}).Preload("CVs").Find(&users)
	return users
}

func IfAppliedToVacancy(vacancyId, userId int) bool {
	var user models.User

	db.DB.Model(&models.User{}).Preload("CVs.Applies").Find(&user, userId)
	if user.Email == "" {
		return false
	}

	for _, cv := range user.CVs {
		for _, applie := range cv.Applies {
			if applie.ID == uint(vacancyId) {
				return true
			}
		}
	}

	return false
}
