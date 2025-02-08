package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pisarraguido/api-notas-Go/controllers"
)

func NotasRoutes(r *gin.Engine) {
	notas := r.Group("/notas")
	{
		notas.GET("/", controllers.ObtenerNotas)
		notas.POST("/", controllers.CrearNota)
		notas.PUT("/:id", controllers.EditarNota)
		notas.DELETE("/:id", controllers.EliminarNota)
	}
}
