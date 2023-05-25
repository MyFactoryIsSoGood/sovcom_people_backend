package models

import "time"

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	CVs      []CV   `json:"cvs"`
}

type Vacancy struct {
	Id          uint              `json:"id" gorm:"primaryKey"`
	Title       string            `json:"title"`
	Company     string            `json:"company"`
	Description string            `json:"description"`
	Templates   []VacancyTemplate `json:"templates"`
	Status      int               `json:"status"` //Поиск, Собес, Новая
	Applies     []Apply           `json:"applies"`
}

type Apply struct {
	Id        uint `json:"id"`
	VacancyId uint `json:"vacancy_id"`
	CVId      uint `json:"cv_id"`
	Status    uint `json:"status"` // Отказ, Приглашение, На рассмотрении
}

type VacancyTemplate struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CV struct {
	Id      uint    `json:"id"`
	Title   string  `json:"title"`
	About   string  `json:"about"`
	UserID  uint    `json:"user_id"`
	Applies []Apply `json:"applies"`
}

type CVTemplate struct {
	Id          uint         `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Blocks      []Experience `json:"blocks"`
}

type Experience struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	DateFrom    time.Time `json:"date_from"`
	DateTo      time.Time `json:"date_to"`
}
