package controllers

import (
	"ApiRestAct1/src/students/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetStudentAgeController struct {
	useCase *application.GetAge
}

func NewGetStudentAgeController(useCase *application.GetAge) *GetStudentAgeController {
	return &GetStudentAgeController{useCase: useCase}
}

func (ctrl *GetStudentAgeController) Execute(c *gin.Context) {
	students, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los estudiantes de 19 años o más"})
		return
	}
	c.JSON(http.StatusOK, students)
}
