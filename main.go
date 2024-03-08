package main

import (
	asignatureInfra "ApiRestAct1/src/asignatures/infraestructure"
	asignatureRoutes "ApiRestAct1/src/asignatures/infraestructure/routes"
	studentInfra "ApiRestAct1/src/students/infraestructure"
	studentRoutes "ApiRestAct1/src/students/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	studentDependencies := studentInfra.InitStudent()
	studentRoutes.ConfigureRoutes(r,
		studentDependencies.CreateStudentController,
		studentDependencies.ListStudentController,
		studentDependencies.ListStudentByIDController,
		studentDependencies.UpdateStudentController,
		studentDependencies.DeleteStudentController,
		studentDependencies.GetStudentByAgeController,
	)

	asignatureDependencies := asignatureInfra.InitAsignature()
	asignatureRoutes.ConfigureRoutesAsignature(
		r,
		asignatureDependencies.CreateAsignatureController,
		asignatureDependencies.ListAsignatureController,
		asignatureDependencies.ListAsignatureByIDController,
		asignatureDependencies.UpdateAsignatureController,
		asignatureDependencies.DeleteAsignatureController,
	)

	r.Run(":8087")
}
