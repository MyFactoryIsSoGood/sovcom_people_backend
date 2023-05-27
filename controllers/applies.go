package controllers

import (
	"awesomeProject/db"
	"awesomeProject/driver"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func PostApply(c *gin.Context) {
//	//
//}

type CreateTestRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Questions   []models.Question `json:"questions"`
}

type CreateCVRequest struct {
	Title  string              `json:"title"`
	About  string              `json:"about"`
	Blocks []models.CVTemplate `json:"blocks"`
}

type CreateApplyRequest struct {
	VacancyID uint   `json:"vacancyId"`
	CVID      uint   `json:"cvId"`
	Comment   string `json:"comment"`
}

func PostTestTask(c *gin.Context) {
	var request CreateTestRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTest := models.Test{
		Title:       request.Title,
		Description: request.Description,
		Questions:   request.Questions,
	}

	// Save the new Test to the database
	err = db.DB.Create(&newTest).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create test"})
		return
	}

	var resp models.Test
	db.DB.Model(&models.Test{}).Preload("Questions.Variants").Find(&resp, newTest.ID)
	c.JSON(http.StatusOK, resp)
}

func PostCV(c *gin.Context) {
	var request CreateCVRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new CV instance
	newCV := models.CV{
		UserID: uint(c.GetInt("userID")),
		Title:  request.Title,
		About:  request.About,
		Blocks: request.Blocks,
	}

	// Save the new CV to the database
	err = db.DB.Create(&newCV).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create CV"})
		return
	}

	var resp models.CV
	db.DB.Model(&models.CV{}).Preload("Blocks.Strokes").Find(&resp, newCV.ID)
	c.JSON(http.StatusOK, resp)
}

func PostApply(c *gin.Context) {
	var request CreateApplyRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new Apply instance
	newApply := models.Apply{
		VacancyId: request.VacancyID,
		CVId:      request.CVID,
		Comment:   request.Comment,
		Status:    models.Wait,
	}

	err = db.DB.Create(&newApply).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create apply"})
		return
	}

	var resp models.Apply
	db.DB.Model(&models.Apply{}).Preload("Stages").Find(&resp, newApply.ID)
	c.JSON(http.StatusOK, resp)
}

func GetCVById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cv *models.CV
	found, cv := driver.GetCVById(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, &cv)
}
