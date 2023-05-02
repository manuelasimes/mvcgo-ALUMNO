package services

import (
	studentCliente "mvc-goALUMNO/clients/student"
	subjectCliente "mvc-goALUMNO/clients/subject"
	tutorCliente "mvc-goALUMNO/clients/tutor"
	"mvc-goALUMNO/utils/errors"
	//falta subject cliente
	"mvc-goALUMNO/dto"
	"mvc-goALUMNO/models"
	e "mvc-goALUMNO/utils/errors"
)

type studentService struct{}

type studentServiceInterface interface {
	GetStudentById(id int) (dto.StudentDto, e.ApiError)
	GetStudents() (dto.StudentsDto, e.ApiError)
	InsertStudent(studentDto dto.StudentDto) (dto.StudentDto, e.ApiError)
	InsertStudentSubject(subjectDto dto.SubjectDto) (dto.StudentDto, errors.ApiError)
}

var (
	StudentService studentServiceInterface
)

func init() {
	StudentService = &studentService{}
}

func (s *studentService) GetStudentById(id int) (dto.StudentDto, e.ApiError) {

	var student models.Student = studentCliente.GetStudentById(id)
	var studentDto dto.StudentDto

	if student.Id == 0 {
		return studentDto, e.NewBadRequestApiError("user not found")
	}
	studentDto.Name = student.Name
	studentDto.LastName = student.LastName
	studentDto.Id = student.Id
	studentDto.Dni = student.Dni
	studentDto.TutorName = student.Tutor.NameTutor
	studentDto.TutorLastName = student.Tutor.LastNameTutor
	for _, subject := range student.Subjects {
		var dtoSubject dto.SubjectDto

		dtoSubject.Name = subject.Name

		studentDto.SubjectsDto = append(studentDto.SubjectsDto, dtoSubject)
	}

	return studentDto, nil
}

func (s *studentService) GetStudents() (dto.StudentsDto, e.ApiError) {

	var students models.Students = studentCliente.GetStudents()
	var studentsDto dto.StudentsDto

	for _, student := range students {
		var studentDto dto.StudentDto

		studentDto.Name = student.Name
		studentDto.LastName = student.LastName
		studentDto.Id = student.Id
		studentDto.Dni = student.Dni
		studentDto.TutorName = student.Tutor.NameTutor
		studentDto.TutorLastName = student.Tutor.LastNameTutor
		studentsDto = append(studentsDto, studentDto)
	}

	return studentsDto, nil
}

func (s *studentService) InsertStudent(studentDto dto.StudentDto) (dto.StudentDto, e.ApiError) {

	var student models.Student
	var tutor models.Tutor

	student.Name = studentDto.Name
	student.LastName = studentDto.LastName
	student.Dni = studentDto.Dni

	tutor.NameTutor = studentDto.TutorName
	tutor.LastNameTutor = studentDto.TutorLastName
	tutor = tutorCliente.InsertTutor(tutor)

	student.Tutor = tutor
	student = studentCliente.InsertStudent(student)

	studentDto.Id = student.Id

	return studentDto, nil
}

func (s *studentService) InsertStudentSubject(subjectDto dto.SubjectDto) (dto.StudentDto, e.ApiError) {
	var subject models.Subject
	subject.Name = subjectDto.Name
	//adding
	subject = subjectCliente.InsertSubject(subject)

	//find student
	var student models.Student = studentCliente.GetStudentById(subjectDto.StudentId)
	var studentDto dto.StudentDto

	studentDto.Name = student.Name
	studentDto.LastName = student.LastName
	studentDto.Id = student.Id
	studentDto.TutorName = student.Tutor.NameTutor
	studentDto.TutorLastName = student.Tutor.LastNameTutor

	for _, subject := range student.Subjects {
		var dtoSubject dto.SubjectDto

		dtoSubject.Name = subject.Name

		studentDto.SubjectsDto = append(studentDto.SubjectsDto, dtoSubject)

	}
	return studentDto, nil
}
