package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
	"ApiRestAct1/src/students/infraestructure/database"
)

type ListStudent struct {
	db domain.IStudent
}

func NewListStudent(db *database.MySQL) *ListStudent {
	return &ListStudent{db: db}
}

func (ls *ListStudent) Execute() ([]entities.Student, error) {
	return ls.db.GetAll()
}
