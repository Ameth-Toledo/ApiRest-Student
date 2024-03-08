package repositories

import "ApiRestAct1/src/asignatures/domain/entities"

type IMessageService interface {
	PublishEvent(eventType string, asignature entities.Asignature) error
}
