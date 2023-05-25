package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const Applicant = 0
const Recruiter = 1
const Customer = 2

const TestType = 0
const CallType = 1

const Reject = 0
const Invite = 1
const Wait = 2

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Role     uint   `json:"role"` // 0-кандидат 1-рекрутер 2-заказчик
	Phone    string `json:"phone"`
	Password string `json:"password"`
	CVs      []CV   `json:"cvs"`
}

type Vacancy struct {
	gorm.Model
	Title       string            `json:"title"`
	Company     string            `json:"company"`
	Description string            `json:"description"`
	Templates   []VacancyTemplate `json:"templates"`
	Status      int               `json:"status"` //Поиск, Собес, Новая
	Applies     []Apply           `json:"applies"`
}

type Apply struct {
	gorm.Model
	VacancyId uint   `json:"vacancy_id"`
	CVId      uint   `json:"cv_id"`
	Comment   string `json:"comment"`
	Status    uint   `json:"status"` // Отказ, Приглашение, На рассмотрении
	Stages    []Stage
}

type Stage struct {
	gorm.Model
	Type   uint `json:"type"` // 0-test 1-call
	Rating uint `json:"rating"`
	Test   Test
	Call   Call
}

type Call struct {
}

type Test struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Questions   []Question
}

type Question struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Variants    []QuestionVariant
	Answer      QuestionVariant
}

type QuestionVariant struct {
	gorm.Model
	Text string `json:"text"`
}

type VacancyTemplate struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CV struct {
	gorm.Model
	Title   string  `json:"title"`
	About   string  `json:"about"`
	UserID  uint    `json:"user_id"`
	Applies []Apply `json:"applies"`
}

type CVTemplate struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Blocks      []Experience `json:"blocks"`
}

type Experience struct {
	gorm.Model
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	DateFrom    time.Time `json:"date_from"`
	DateTo      time.Time `json:"date_to"`
}
