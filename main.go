package main

import (
	"awesomeProject/controllers"
	"awesomeProject/db"
	"awesomeProject/driver"
	"awesomeProject/mail"
	"awesomeProject/middleware"
	"awesomeProject/models"
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
	app.Use(middleware.CORSMiddleware())

	auth := app.Group("/auth")
	api := app.Group("/api", middleware.AuthMiddleware())

	auth.POST("/signIn", controllers.Auth)
	auth.POST("/signUp", controllers.SignUp)

	api.GET("/users/:id", controllers.GetUserById)
	api.GET("/users", controllers.GetAllUsers)

	api.POST("/vacancies", controllers.PostVacancy)
	api.GET("/vacancies", controllers.GetAllVacancies)
	api.GET("/vacancies/:id", controllers.GetVacancyById)

	api.POST("/cv", controllers.PostCV)

	api.POST("/applies", controllers.PostApply)
	//api.GET("/vacancy/:id/applies", controllers.GetVacancyApplies)
	api.POST("/testTask", controllers.PostTestTask)

	err := app.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
