package controllers

import (
	"awesomeProject/driver"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *models.User
	found, user := driver.GetUserById(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	users = driver.GetUsers()

	c.JSON(http.StatusOK, users)
}

func IfAppliedToVacancy(c *gin.Context) {
	vacancyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt("userID")

	appliedToVacancy := driver.IfAppliedToVacancy(vacancyId, userId)
	c.JSON(http.StatusOK, gin.H{"applied": appliedToVacancy})
}
