package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"github.com/jinzhu/gorm"
)

type FullCV struct {
	gorm.Model
	Title  string              `json:"title"`
	About  string              `json:"about"`
	User   models.User         `json:"user"`
	Blocks []models.CVTemplate `json:"blocks"`
}

type FullApply struct {
	gorm.Model
	CV      FullCV         `json:"cv"`
	Comment string         `json:"comment"`
	Status  uint           `json:"status"`
	Stages  []models.Stage `json:"stages"`
}

type VacancyResponse struct {
	Title       string                   `json:"title"`
	Company     string                   `json:"company"`
	Description string                   `json:"description"`
	Templates   []models.VacancyTemplate `json:"templates"`
	Status      uint                     `json:"status"`
	Applies     []FullApply              `json:"applies"`
}

func GetVacancyById(id int) (bool, *VacancyResponse) {
	var vacancy models.Vacancy
	var fullApplies []FullApply

	db.DB.Model(&models.Vacancy{}).Preload("Templates").Preload("Applies.Stages").Find(&vacancy, id)
	for _, apply := range vacancy.Applies {
		var cv *models.CV
		var user *models.User

		_, cv = GetCVById(int(apply.CVId))
		_, user = GetUserById(int(cv.UserID))
		fullCv := FullCV{Model: gorm.Model{ID: apply.ID,
			CreatedAt: apply.CreatedAt,
		}, Title: cv.Title, About: cv.About, User: *user, Blocks: cv.Blocks}
		fullApplies = append(fullApplies, FullApply{
			Model: gorm.Model{ID: apply.ID,
				CreatedAt: apply.CreatedAt,
			},
			CV:      fullCv,
			Comment: apply.Comment,
			Status:  apply.Status,
			Stages:  apply.Stages,
		})
	}

	if vacancy.Title != "" {
		resp := VacancyResponse{
			Title:       vacancy.Title,
			Company:     vacancy.Company,
			Description: vacancy.Description,
			Templates:   vacancy.Templates,
			Status:      vacancy.Status,
			Applies:     fullApplies,
		}
		return true, &resp
	}

	return false, &VacancyResponse{}
}

func CreateVacancy(vacancy models.Vacancy, templates []models.VacancyTemplate) (error, models.Vacancy) {
	err := db.DB.Create(&vacancy).Error
	if err != nil {
		return err, models.Vacancy{}
	}

	for _, template := range templates {
		newTemplate := models.VacancyTemplate{
			VacancyId:   vacancy.ID,
			Title:       template.Title,
			Description: template.Description,
		}

		db.DB.Model(&models.VacancyTemplate{}).Create(&newTemplate)
	}

	db.DB.Model(&models.Vacancy{}).Preload("Templates").First(&vacancy, vacancy.ID)
	return nil, vacancy
}
