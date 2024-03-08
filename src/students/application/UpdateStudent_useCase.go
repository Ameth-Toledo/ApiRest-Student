package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
	"ApiRestAct1/src/students/infraestructure/database"
)

type UpdateStudent struct {
	db domain.IStudent
}

func NewUpdateStudent(db *database.MySQL) *UpdateStudent {
	return &UpdateStudent{db: db}
}

func (us *UpdateStudent) Execute(student entities.Student) error {
	return us.db.Update(student)
}
