package routes

import (
	"github.com/gin-gonic/gin"
	"api-notas-Go/controllers"
)

func AuthRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.POST("/register", controllers.Registro)
		routes.POST("/login", controllers.Login)
	}
}
