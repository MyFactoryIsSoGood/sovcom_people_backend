package controllers

import (
	"awesomeProject/driver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthBody struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Auth(c *gin.Context) {
	body := AuthBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	found, user := driver.CheckCredentials(body.Phone, body.Password)
	if found {
		c.JSON(http.StatusOK, user)
		return
	} else {
		c.Status(http.StatusUnauthorized)
		return
	}
}
