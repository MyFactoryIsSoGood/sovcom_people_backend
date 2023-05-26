package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
)

func GetVacancyById(id int) (bool, *models.Vacancy) {
	var vacancy models.Vacancy

	db.DB.Model(&models.Vacancy{}).Where("id = ?", id).Scan(&vacancy)
	if vacancy.Title != "" {
		return true, &vacancy
	}

	return false, &models.Vacancy{}
}

func CreateVacancy(vacancy models.Vacancy, templates []models.VacancyTemplate) (error, models.Vacancy) {
	err := db.DB.Create(&vacancy).Error
	if err != nil {
		return err, models.Vacancy{}
	}

	for _, template := range templates {
		newTemplate := models.VacancyTemplate{
			VacancyId:   vacancy.ID,
			Title:       template.Title,
			Description: template.Description,
		}

		db.DB.Model(&models.VacancyTemplate{}).Create(&newTemplate)
	}

	db.DB.Model(&models.Vacancy{}).Preload("Templates").First(&vacancy, vacancy.ID)
	return nil, vacancy
}
