package application

import (
	"ApiRestAct1/src/asignatures/application/repositories"
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
	"log"
)

type CreateAsignature struct {
	asignatureRepo      domain.IAsignature
	serviceNotification repositories.IMessageService
}

func NewCreateAsignature(asignatureRepo domain.IAsignature, serviceNotification repositories.IMessageService) *CreateAsignature {
	return &CreateAsignature{
		asignatureRepo:      asignatureRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *CreateAsignature) Execute(asignature entities.Asignature) (entities.Asignature, error) {
	log.Println("💾 Guardando asignatura en la base de datos...")

	created, err := c.asignatureRepo.Save(asignature)
	if err != nil {
		log.Println("❌ Error al guardar en la base de datos:", err)
		return entities.Asignature{}, err
	}

	log.Println("✅ Asignatura guardada correctamente. Publicando evento...")

	err = c.serviceNotification.PublishEvent("AsignatureCreated", created)
	if err != nil {
		log.Println("❌ Error al notificar la creación de la asignatura:", err)
		return entities.Asignature{}, err
	}

	log.Println("🎉 Asignatura creada y notificada correctamente.")
	return created, nil
}
