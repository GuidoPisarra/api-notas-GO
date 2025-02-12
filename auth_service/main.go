package main

import (
	"auth_service/config"
	"auth_service/database"
	"auth_service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectMongoDB()

	config.SetupLogging()

	config.Init()

	r := gin.Default()

	// Cargar las rutas de autenticación
	routes.SetupAuthRoutes(r)

	port := config.GetEnv("AUTH_PORT", "8081") // Usa 8081 para autenticación
	log.Println("Auth Service corriendo en el puerto", port)
	r.Run(":" + port)
}
