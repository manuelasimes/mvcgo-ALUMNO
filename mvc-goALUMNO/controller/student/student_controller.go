package studentController

import (
	"mvc-goALUMNO/dto"
	service "mvc-goALUMNO/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetStudentById(c *gin.Context) {
	log.Debug("Student id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var studentDto dto.StudentDto

	studentDto, err := service.StudentService.GetStudentById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, studentDto)
}

func GetStudents(c *gin.Context) {
	var studentsDto dto.StudentsDto
	studentsDto, err := service.StudentService.GetStudents()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, studentsDto)
}

func StudentInsert(c *gin.Context) {
	var studentDto dto.StudentDto
	err := c.BindJSON(&studentDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	studentDto, er := service.StudentService.InsertStudent(studentDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, studentDto)
}
func InsertStudentSubject(c *gin.Context) {

	log.Debug("Adding Subject to student: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var subjectDto dto.SubjectDto
	err := c.BindJSON(&subjectDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	subjectDto.StudentId = id

	var studentDto dto.StudentDto

	studentDto, er := service.StudentService.InsertStudentSubject(subjectDto)
	//Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, studentDto)
}
