package db

import (
	"awesomeProject/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var DB *gorm.DB

func Connect() error {
	//Docker config:
	args := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable host=%v port=5432",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"))

	db, err := gorm.Open("postgres", args)

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{}, &models.Vacancy{}, &models.CV{}, &models.Experience{}, &models.VacancyTemplate{}, &models.CVTemplate{}, &models.Apply{})
	DB = db
	return nil
}
