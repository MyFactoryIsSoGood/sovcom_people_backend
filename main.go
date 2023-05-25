package main

import (
	"awesomeProject/controllers"
	"awesomeProject/db"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

// Initialize инициализирует все необходимые подключения и зависимости
func Initialize() {
	var initErrors []string

	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			initErrors = append(initErrors, err.Error())
		}
	}

	err := db.Connect()
	if err != nil {
		initErrors = append(initErrors, err.Error())
	}

	if len(initErrors) != 0 {
		panic(fmt.Sprintf("Запуск приложения невозможен из-за следующих ошибок инициализации:\n %s", strings.Join(initErrors, ",\n")))
	}
}

func Mock() {
	serzh := models.User{
		FullName: "Serzh Galeta",
		Role:     models.Applicant,
		Email:    "galeta@mail.ru",
		Phone:    "123",
		Password: "123",
	}

	db.DB.Model(&models.User{}).Create(&serzh)
}

func main() {
	Initialize()
	Mock()

	app := gin.Default()
	app.Use(cors.Default())

	app.GET("/", func(context *gin.Context) {
		context.Status(200)
	})
	app.POST("/auth", controllers.Auth)

	err := app.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
