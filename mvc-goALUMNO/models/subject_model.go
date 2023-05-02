package models

type Subject struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(350);not full"`

	Student Students `gorm:"many2many:alumno_materia"`
}

type Subjects []Subject
