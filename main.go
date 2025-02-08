package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pisarraguido/api-notas-Go/database" 
	"github.com/pisarraguido/api-notas-Go/routes"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
)

// Estructura de Nota
type Nota struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Titulo   string `json:"titulo" bson:"titulo"`
	Contenido string `json:"contenido" bson:"contenido"`
}

func main() {
	// Conectar a la base de datos
	database.ConnectMongoDB()

	// Crear router
	r := gin.Default()

	// Cargar rutas desde la carpeta routes
	routes.NotasRoutes(r)

	// Iniciar servidor
	r.Run(":8080")
}

// Obtener todas las notas
func obtenerNotas(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notas []Nota
	collection := database.GetCollection("notas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las notas"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var nota Nota
		cursor.Decode(&nota)
		notas = append(notas, nota)
	}

	c.JSON(http.StatusOK, notas)
}

// Crear una nueva nota
func crearNota(c *gin.Context) {
	var nuevaNota Nota
	if err := c.BindJSON(&nuevaNota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err := collection.InsertOne(ctx, nuevaNota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la nota"})
		return
	}

	c.JSON(http.StatusCreated, nuevaNota)
}
