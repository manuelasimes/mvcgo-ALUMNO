package models

type Tutor struct {
	Id            int    `gorm:"primary_key"`
	NameTutor     string `gorm:"type:varchar(350);not full"`
	LastNameTutor string `gorm:"type:varchar(350);not full"`
}

type Tutors []Tutor
