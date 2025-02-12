package routes

import (
	"auth_service/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes define las rutas de autenticación
func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Registro)
	}
}
