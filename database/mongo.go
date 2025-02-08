package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"api-notas-Go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// Conectar a MongoDB usando las variables de entorno
func ConnectMongoDB() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Construcción de la URI de conexión a MongoDB
	var mongoURI string
	if cfg.MongoUser != "" && cfg.MongoPassword != "" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", cfg.MongoHost, cfg.MongoPort)
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	fmt.Println("✅ Conectado a MongoDB en", mongoURI)
	MongoClient = client
	MongoDB = client.Database(cfg.MongoDB)
}

// Obtener una colección
func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Collection(collectionName)
}
