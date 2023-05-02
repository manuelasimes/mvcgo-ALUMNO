package student

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"mvc-goALUMNO/models"
)

var Db *gorm.DB

func GetStudentById(id int) models.Student {
	var student models.Student

	Db.Where("id = ?", id).Preload("Tutor").Preload("Subjects").First(&student)
	log.Debug("Student: ", student)

	return student
}

func GetStudents() models.Students {
	var students models.Students
	Db.Preload("Tutor").Find(&students)

	log.Debug("Students: ", students)

	return students
}

func InsertStudent(student models.Student) models.Student {
	result := Db.Create(&student)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Student Inserted: ", student.Id)
	return student
}
