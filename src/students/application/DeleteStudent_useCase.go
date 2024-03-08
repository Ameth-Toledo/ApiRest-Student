package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/infraestructure/database"
)

type DeleteStudent struct {
	db domain.IStudent
}

func NewDeleteStudent(db *database.MySQL) *DeleteStudent {
	return &DeleteStudent{db: db}
}

func (ds *DeleteStudent) Execute(id int) error {
	return ds.db.Delete(id)
}
