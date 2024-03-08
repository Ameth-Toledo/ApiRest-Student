package infraestructure

import (
	"ApiRestAct1/src/asignatures/application"
	"ApiRestAct1/src/asignatures/infraestructure/adapter"
	"ApiRestAct1/src/asignatures/infraestructure/controllers"
	"ApiRestAct1/src/asignatures/infraestructure/database"
	"log"
)

type DependenciesAsignature struct {
	CreateAsignatureController   *controllers.CreateAsignatureController
	ListAsignatureController     *controllers.ListAsignatureController
	ListAsignatureByIDController *controllers.ListAsignatureByIdController
	UpdateAsignatureController   *controllers.UpdateAsignatureController
	DeleteAsignatureController   *controllers.DeleteAsignatureController
}

func InitAsignature() *DependenciesAsignature {
	ps := database.NewMySQL()

	rmqClient, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	return &DependenciesAsignature{
		CreateAsignatureController:   controllers.NewCreateAsignatureController(application.NewCreateAsignature(ps, rmqClient), ps), // Pass both parameters
		ListAsignatureController:     controllers.NewListAsignatureController(application.NewListAsignature(ps)),
		ListAsignatureByIDController: controllers.NewListAsignatureByIdController(application.NewListAsignatureById(ps)),
		UpdateAsignatureController:   controllers.NewUpdateAsignatureController(application.NewUpdateAsignature(ps)),
		DeleteAsignatureController:   controllers.NewDeleteAsignatureController(application.NewDeleteAsignature(ps)),
	}
}
