package routes

import (
	"github.com/gin-gonic/gin"
	"api-notas-Go/controllers"
	"api-notas-Go/middlewares"
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
