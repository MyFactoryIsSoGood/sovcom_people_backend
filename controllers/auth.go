package controllers

import (
	"awesomeProject/driver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpBody struct {
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
	Location string `json:"location"`
	WorkMode string `json:"workMode"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     uint   `json:"role" `
}

func Auth(c *gin.Context) {
	body := AuthBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	found, user := driver.CheckCredentials(body.Email, body.Password)
	if found {
		c.JSON(http.StatusOK, user)
		return
	}
	c.Status(http.StatusUnauthorized)
	return
}

func SignUp(c *gin.Context) {
	body := SignUpBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, user := driver.CreateUser(body.Phone, body.Email, body.Password, body.FullName, body.Location, body.WorkMode, body.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
