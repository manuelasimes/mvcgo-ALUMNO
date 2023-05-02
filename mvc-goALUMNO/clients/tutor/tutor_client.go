package tutor

import (
	"mvc-goALUMNO/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetTutorById(id int) models.Tutor {
	var tutor models.Tutor

	Db.Where("id = ?", id).Preload("Subject").Preload("Students").First(&tutor)
	log.Debug("Tutor: ", tutor)

	return tutor
}

func InsertTutor(tutor models.Tutor) models.Tutor {
	result := Db.Create(&tutor)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Tutor Inserted: ", tutor.Id)
	return tutor
}
