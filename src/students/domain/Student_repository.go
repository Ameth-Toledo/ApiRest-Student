package domain

import "ApiRestAct1/src/students/domain/entities"

type IStudent interface {
	Save(student entities.Student) error
	GetAll() ([]entities.Student, error)
	GetById(id int) (entities.Student, error)
	Update(student entities.Student) error
	Delete(id int) error
	GetByAge(minAge int) ([]entities.Student, error)
}
