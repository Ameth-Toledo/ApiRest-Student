package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
	"ApiRestAct1/src/students/infraestructure/database"
)

type CreateStudent struct {
	db domain.IStudent
}

func NewCreateStudent(db *database.MySQL) *CreateStudent {
	return &CreateStudent{db: db}
}

func (cp *CreateStudent) Execute(student entities.Student) error {
	return cp.db.Save(student)
}
