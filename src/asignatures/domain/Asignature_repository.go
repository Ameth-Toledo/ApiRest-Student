package domain

import "ApiRestAct1/src/asignatures/domain/entities"

type IAsignature interface {
	Save(asignature entities.Asignature) (entities.Asignature, error)
	GetAll() ([]entities.Asignature, error)
	GetById(id int) (entities.Asignature, error)
	Update(asignature entities.Asignature) error
	Delete(id int) error
}
