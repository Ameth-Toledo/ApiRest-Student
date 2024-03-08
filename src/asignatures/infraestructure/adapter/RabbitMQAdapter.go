package adapter

import (
	"ApiRestAct1/src/asignatures/application/repositories"
	"ApiRestAct1/src/asignatures/domain/entities"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repositories.IMessageService = (*RabbitMQAdapter)(nil)

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("amqp://toledo:12345@35.170.134.124:5672/")
	if err != nil {
		log.Println("Error al conectar con RabbitMQ:", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error al abrir el canal:", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"asignatures",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error al declarar la cola:", err)
		return nil, err
	}

	log.Println("Conexión establecida con RabbitMQ y cola 'asignatures' declarada.")

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, asignature entities.Asignature) error {
	body, err := json.Marshal(asignature)
	if err != nil {
		log.Println("Error al convertir a JSON:", err)
		return err
	}

	log.Printf("Publicando mensaje en la cola 'asignatures': %s\n", body)

	err = r.ch.Publish(
		"",
		"asignatures",
		true,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		log.Println("Error al enviar mensaje a RabbitMQ:", err)
		return err
	}

	log.Println("Mensaje publicado exitosamente en RabbitMQ")
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Error cerrando canal de RabbitMQ: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error cerrando conexión de RabbitMQ: %v", err)
	}
}
