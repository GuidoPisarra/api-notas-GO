package controllers

import (
	"context"
	"log"
	"net/http"
	"notas_service/database"
	"notas_service/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObtenerNotas godoc
// @Summary Obtiene todas las notas
// @Description Obtiene una lista de todas las notas almacenadas
// @Tags Notas
// @Accept json
// @Produce json
// @Success 200 {array} models.Nota
// @Failure 500 {object} map[string]interface{}
// @Router /notas [get]
func ObtenerNotas(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notas []models.Nota
	collection := database.GetCollection("notas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al obtener notas: %v\n", err) // Log del error
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

// CrearNota godoc
// @Summary Crea una nueva nota
// @Description Crea una nueva nota y la guarda en la base de datos
// @Tags Notas
// @Accept json
// @Produce json
// @Param nota body models.Nota true "Nueva nota"
// @Success 201 {object} models.Nota
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notas [post]
func CrearNota(c *gin.Context) {
	var nuevaNota models.Nota
	if err := c.BindJSON(&nuevaNota); err != nil {
		log.Printf("Error al vincular datos: %v\n", err) // Log de error de validación
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err := collection.InsertOne(ctx, nuevaNota)
	if err != nil {
		log.Println("Error al crear la nota: ", err) // Log del error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la nota"})
		return
	}

	c.JSON(http.StatusCreated, nuevaNota)
}

// EditarNota godoc
// @Summary Edita una nota existente
// @Description Actualiza los detalles de una nota específica en la base de datos
// @Tags Notas
// @Accept json
// @Produce json
// @Param id path string true "ID de la nota"
// @Param nota body models.Nota true "Nota actualizada"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notas/{id} [put]
func EditarNota(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("ID inválido proporcionado: %v\n", err) // Log de error de ID inválido
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var nota models.Nota
	if err := c.BindJSON(&nota); err != nil {
		log.Printf("Error al vincular datos: %v\n", err) // Log de error de validación
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": nota})
	if err != nil {
		log.Printf("Error al actualizar la nota: %v\n", err) // Log del error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la nota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nota actualizada correctamente"})
}

// EliminarNota godoc
// @Summary Elimina una nota
// @Description Elimina una nota de la base de datos según el ID proporcionado
// @Tags Notas
// @Accept json
// @Produce json
// @Param id path string true "ID de la nota"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notas/{id} [delete]
func EliminarNota(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("ID inválido proporcionado: %v\n", err) // Log de error de ID inválido
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("notas")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Printf("Error al eliminar la nota: %v\n", err) // Log del error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la nota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nota eliminada correctamente"})
}
