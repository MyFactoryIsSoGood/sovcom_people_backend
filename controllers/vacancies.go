package controllers

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateVacancyBody struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Templates   []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"templates"`
}

func GetAllVacancies(c *gin.Context) {
	var vacancies []models.Vacancy

	// Retrieve all vacancies from the database and preload associated templates and applies
	err := db.DB.Preload("Templates").Find(&vacancies).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vacancies"})
		return
	}

	c.JSON(http.StatusOK, vacancies)
}

func PostVacancy(c *gin.Context) {
	var vacancy CreateVacancyBody
	err := c.ShouldBindJSON(&vacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new Vacancy instance
	newVacancy := models.Vacancy{
		Title:       vacancy.Title,
		Company:     vacancy.Company,
		Description: vacancy.Description,
		Status:      models.Searching,
	}

	// Save the new vacancy to the database
	err = db.DB.Create(&newVacancy).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vacancy"})
		return
	}
	for _, template := range vacancy.Templates {
		newTemplate := models.VacancyTemplate{
			Title:       template.Title,
			Description: template.Description,
		}
		err = db.DB.Create(&newTemplate).Error
		db.DB.Model(&newTemplate).Association("Vacancy").Append(newVacancy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create template"})
			return
		}
	}

	c.JSON(http.StatusOK, newVacancy)
}
