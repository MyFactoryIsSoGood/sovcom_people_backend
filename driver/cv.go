package driver

import (
	"awesomeProject/db"
	"awesomeProject/models"
)

func GetCVById(id int) (bool, *models.CV) {
	var cv models.CV

	db.DB.Model(&models.CV{}).Preload("Blocks.Strokes").Find(&cv, id)
	if cv.Title != "" {
		return true, &cv
	}

	return false, &models.CV{}
}
