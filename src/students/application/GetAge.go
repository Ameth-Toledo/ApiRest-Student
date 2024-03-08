package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
	"ApiRestAct1/src/students/infraestructure/database"
)

type GetAge struct {
	db domain.IStudent
}

func NewGetAge(db *database.MySQL) *GetAge {
	return &GetAge{db: db}
}

func (ga *GetAge) Execute() ([]entities.Student, error) {
	return ga.db.GetByAge(19)
}
