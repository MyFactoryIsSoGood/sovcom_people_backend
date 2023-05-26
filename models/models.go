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

const Searching = 0
const Interview = 1
const New = 2
const Closed = 3

const Reject = 0
const Invite = 1
const Wait = 2

type User struct {
	gorm.Model
	FullName string `json:"fullName"`
	Role     uint   `json:"role"` // 0-кандидат 1-рекрутер 2-заказчик
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CVs      []CV   `json:"cvs"`
}

type Vacancy struct {
	gorm.Model
	Title       string            `json:"title"`
	Company     string            `json:"company"`
	Description string            `json:"description"`
	Templates   []VacancyTemplate `json:"templates"`
	Status      uint              `json:"status"` //Поиск, Собес, Новая
	Applies     []Apply           `json:"applies"`
}

type Apply struct {
	gorm.Model
	VacancyId uint    `json:"vacancyId"`
	CVId      uint    `json:"cvId"`
	Comment   string  `json:"comment"`
	Status    uint    `json:"status"` // Отказ, Приглашение, На рассмотрении
	Stages    []Stage `json:"stages"`
}

type Stage struct {
	gorm.Model
	ApplyId uint `json:"applyId"`
	Type    uint `json:"type"` // 0-test 1-call
	Rating  uint `json:"rating"`
	Test    Test `json:"test"`
	Call    Call `json:"call"`
}

type Call struct {
}

type Test struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	gorm.Model
	TestId      uint              `json:"testId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Answer      string            `json:"answer"`
	Variants    []QuestionVariant `json:"variants"`
}

type QuestionVariant struct {
	gorm.Model
	QuestionId uint   `json:"questionId"`
	Text       string `json:"text"`
}

type VacancyTemplate struct {
	gorm.Model
	VacancyId   uint   `json:"vacancyId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CV struct {
	gorm.Model
	Title   string       `json:"title"`
	About   string       `json:"about"`
	UserID  uint         `json:"userId"`
	Blocks  []CVTemplate `json:"blocks"`
	Applies []Apply      `json:"applies"`
}

type CVTemplate struct {
	gorm.Model
	CVId    uint         `json:"cvId"`
	Title   string       `json:"title"`
	Strokes []Experience `json:"strokes"`
}

type Experience struct {
	gorm.Model
	CVTemplateId uint      `json:"CVTemplateId"`
	Title        string    `json:"title"`
	Subtitle     string    `json:"subtitle"`
	Description  string    `json:"description"`
	DateFrom     time.Time `json:"dateFrom"`
	DateTo       time.Time `json:"dateTo"`
}
