package main

import (
	"notas_service/config"
	"notas_service/database"
	"notas_service/routes"

	_ "notas_service/docs" // Commented out as the package does not exist

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Notas
// @version 1.0
// @description API para gestionar notas - Guido Pisarra

// @host localhost:8080
// @BasePath /

func main() {
	// Conectar a la base de datos
	database.ConnectMongoDB()

	config.SetupLogging()

	config.Init()

	router := gin.Default()

	// Ruta para servir la documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//routes.AuthRoutes(router)
	routes.NotasRoutes(router)

	router.Run(":8082")
}
