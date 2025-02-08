package main

import (
	"api-notas-Go/config"
	"api-notas-Go/database"
	"api-notas-Go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	database.ConnectMongoDB()

	config.Init()

	router := gin.Default()

	routes.AuthRoutes(router)
	routes.NotasRoutes(router)

	router.Run(":8080")
}
