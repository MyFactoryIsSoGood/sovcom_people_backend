package main

import (
	"awesomeProject/db"
	"awesomeProject/driver"
	"awesomeProject/mail"
	"awesomeProject/models"
	"awesomeProject/routes"
	"fmt"
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
	if found, _ := driver.GetUserByMail("galeta@mail.ru"); !found {
		db.DB.Model(&models.User{}).Create(&models.User{
			FullName: "Serzh Galeta",
			Role:     models.Applicant,
			Email:    "galeta@mail.ru",
			Phone:    "123",
			Password: "123",
		})
	}
}

func main() {
	mail.SendMail("", "")
	Initialize()
	Mock()

	app := gin.Default()
	routes.InitializeRoutes(app)

	err := app.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
