package app

import (
	log "github.com/sirupsen/logrus"
	studentController "mvc-goALUMNO/controller/student"
)

func mapUrls() {

	// Users Mapping
	router.GET("/student/:id", studentController.GetStudentById)
	router.GET("/student", studentController.GetStudents)
	router.POST("/student", studentController.StudentInsert)
	router.POST("/student/:id/subject", studentController.InsertStudentSubject)

	log.Info("Finishing mappings configurations")
}
