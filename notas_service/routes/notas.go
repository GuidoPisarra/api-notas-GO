package routes

import (
	"notas_service/controllers"
	"notas_service/handlers"
	"notas_service/middlewares"
	"notas_service/repository"
	"notas_service/services"

	"github.com/gin-gonic/gin"
)

func NotasRoutes(router *gin.Engine) {

	routes := router.Group("/notas")
	routes.Use(middlewares.AuthMiddleware())

	// Inicializar el repositorio, servicio y handler
	notaRepo := &repository.NotaRepository{}
	notaService := services.NuevoNotaService(notaRepo)
	notaHandler := handlers.NuevoNotaHandler(notaService)

	// Definir rutas con el handler

	{
		routes.GET("/", notaHandler.ObtenerNotas)
		routes.POST("/", controllers.CrearNota)
		routes.PUT("/:id", controllers.EditarNota)
		routes.DELETE("/:id", controllers.EliminarNota)
	}
}
