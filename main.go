package main

import (
	"awesomeProject/db"
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

func main() {
	Initialize()

	app := gin.Default()

	err := app.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
