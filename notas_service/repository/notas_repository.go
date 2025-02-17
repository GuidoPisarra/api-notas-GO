package repository

import (
	"context"
	"notas_service/database"
	"notas_service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// NotaRepository estructura para acceder a la DB
type NotaRepository struct{}

// ObtenerNotas obtiene todas las notas de la base de datos
func (r *NotaRepository) ObtenerNotas() ([]models.Nota, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notas []models.Nota
	collection := database.GetCollection("notas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &notas); err != nil {
		return nil, err
	}

	return notas, nil
}
