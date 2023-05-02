package subject

import (
	"mvc-goALUMNO/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetSubjectById(id int) models.Subject {
	var subject models.Subject

	Db.Where("id = ?", id).Preload("Tutor").Preload("Students").First(&subject)
	log.Debug("Subject: ", subject)

	return subject
}

func InsertSubject(subject models.Subject) models.Subject {
	result := Db.Create(&subject)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Subject Inserted: ", subject.Id)
	return subject
}
