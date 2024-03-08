package infraestructure

import (
	"ApiRestAct1/src/students/application"
	"ApiRestAct1/src/students/infraestructure/controllers"
	"ApiRestAct1/src/students/infraestructure/database"
)

type DependenciesStudent struct {
	CreateStudentController   *controllers.CreateStudentController
	ListStudentController     *controllers.ListStudentController
	ListStudentByIDController *controllers.ListStudentByIDController
	UpdateStudentController   *controllers.UpdateStudentController
	DeleteStudentController   *controllers.DeleteStudentController
	GetStudentByAgeController *controllers.GetStudentAgeController
}

func InitStudent() *DependenciesStudent {
	ps := database.NewMySQL()

	return &DependenciesStudent{
		CreateStudentController:   controllers.NewCreateStudentController(application.NewCreateStudent(ps)),
		ListStudentController:     controllers.NewListStudentController(application.NewListStudent(ps)),
		ListStudentByIDController: controllers.NewListStudentByIDController(application.NewListStudentById(ps)),
		UpdateStudentController:   controllers.NewUpdateStudentController(application.NewUpdateStudent(ps)),
		DeleteStudentController:   controllers.NewDeleteStudentController(application.NewDeleteStudent(ps)),
		GetStudentByAgeController: controllers.NewGetStudentAgeController(application.NewGetAge(ps)),
	}
}
