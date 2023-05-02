package models

type Student struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	Dni      int    `gorm:"type:int"`

	Tutor   Tutor `gorm:"foreignkey:TutorId"`
	TutorId int

	Subjects Subjects `gorm:"many2many:alumno_materia"`
}
type Students []Student
