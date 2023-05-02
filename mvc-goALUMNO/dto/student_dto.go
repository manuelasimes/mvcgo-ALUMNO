package dto

type StudentDto struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	LastName      string      `json:"last_name"`
	Dni           int         `json:"dni"`
	TutorName     string      `json:"tutor_name"`
	TutorLastName string      `json:"tutor_last_name"`
	SubjectsDto   SubjectsDto `json:"subjects_dto,omitempty"`
}
type StudentsDto []StudentDto
