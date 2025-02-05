package config

import (
	"os"
	"log"
	"fmt"
	"github.com/joho/godotenv"
)

// Config estructura de configuración
type Config struct {
	MongoHost   string
	MongoPort   string
	MongoUser   string
	MongoPassword string
	MongoDB     string
}

// Cargar las variables de entorno y devolver configuración
func LoadConfig() (Config, error) {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	cfg := Config{
		MongoHost:   os.Getenv("MONGO_HOST"),  
		MongoPort:   os.Getenv("MONGO_PORT"),  
		MongoUser:   os.Getenv("MONGO_USER"),  
		MongoPassword: os.Getenv("MONGO_PASSWORD"),  
		MongoDB:     os.Getenv("MONGO_DB"),  
	}

	if cfg.MongoHost == "" || cfg.MongoPort == "" || cfg.MongoDB == "" {
		log.Fatalf("Las variables de entorno no están configuradas correctamente")
		return cfg, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
