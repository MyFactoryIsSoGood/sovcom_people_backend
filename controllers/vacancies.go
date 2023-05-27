package controllers

import (
	"awesomeProject/db"
	"awesomeProject/driver"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateVacancyBody struct {
	Title       string `json:"title" binding:"required"`
	Company     string `json:"company" binding:"required"`
	Description string `json:"description"`
	Templates   []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"templates"`
}

func GetAllVacancies(c *gin.Context) {
	var vacancies []models.Vacancy

	err := db.DB.Model(&models.Vacancy{}).Preload("Templates").Find(&vacancies).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vacancies"})
		return
	}

	c.JSON(http.StatusOK, vacancies)
}

func GetVacancyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vacancy *driver.VacancyResponse
	found, vacancy := driver.GetVacancyById(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, &vacancy)
}

func PostVacancy(c *gin.Context) {
	var vacancy CreateVacancyBody
	err := c.ShouldBindJSON(&vacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newVacancy := models.Vacancy{
		Title:       vacancy.Title,
		Company:     vacancy.Company,
		Description: vacancy.Description,
		Status:      models.Searching,
	}

	var templates []models.VacancyTemplate
	for _, template := range vacancy.Templates {
		newTemplate := models.VacancyTemplate{
			VacancyId:   newVacancy.ID,
			Title:       template.Title,
			Description: template.Description,
		}
		templates = append(templates, newTemplate)
	}

	err, newVacancy = driver.CreateVacancy(newVacancy, templates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newVacancy)
}

func UpdateVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vacancy models.Vacancy
	db.DB.Model(&models.Vacancy{}).Preload("Templates").Find(&vacancy, id)

	var newVacancy CreateVacancyBody
	err = c.ShouldBindJSON(&newVacancy)
	fmt.Println(newVacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, template := range vacancy.Templates {
		db.DB.Model(&models.VacancyTemplate{}).Delete(&template)
	}
	vacancy.Templates = []models.VacancyTemplate{}
	for _, template := range newVacancy.Templates {
		newTemplate := models.VacancyTemplate{
			VacancyId:   vacancy.ID,
			Title:       template.Title,
			Description: template.Description,
		}
		db.DB.Model(&models.VacancyTemplate{}).Create(&newTemplate)
	}

	vacancy.Title = newVacancy.Title
	vacancy.Company = newVacancy.Company
	vacancy.Description = newVacancy.Description
	db.DB.Model(&models.Vacancy{}).Update(&vacancy)

	_, resp := driver.GetVacancyById(id)
	c.JSON(http.StatusOK, resp)
}

func GetVacancyApplies(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vacancy *driver.VacancyResponse
	found, vacancy := driver.GetVacancyById(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, &vacancy.Applies)
}
