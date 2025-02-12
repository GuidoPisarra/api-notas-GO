package routes

import (
	"notas_service/controllers"
	"notas_service/middlewares"

	"github.com/gin-gonic/gin"
)

func NotasRoutes(router *gin.Engine) {
	routes := router.Group("/notas")
	routes.Use(middlewares.AuthMiddleware())
	{
		routes.GET("/", controllers.ObtenerNotas)
		routes.POST("/", controllers.CrearNota)
		routes.PUT("/:id", controllers.EditarNota)
		routes.DELETE("/:id", controllers.EliminarNota)
	}
}
