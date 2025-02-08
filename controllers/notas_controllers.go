package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"api-notas-Go/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"api-notas-Go/models"
)



// Obtener todas las notas
func ObtenerNotas(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notas []models.Nota
	collection := database.GetCollection("notas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las notas"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var nota models.Nota
		cursor.Decode(&nota)
		notas = append(notas, nota)
	}

	c.JSON(http.StatusOK, notas)
}

// Crear una nueva nota
func CrearNota(c *gin.Context) {
	var nuevaNota models.Nota
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

// Editar una nota
func EditarNota(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var nota models.Nota
	if err := c.BindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": nota})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la nota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nota actualizada correctamente"})
}

// Eliminar una nota
func EliminarNota(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la nota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nota eliminada correctamente"})
}