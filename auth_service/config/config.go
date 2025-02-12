package config

import (
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
)

// Config estructura de configuración
type Config struct {
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoPort     string
	MongoDB       string
	JWTSecretKey  string
}

var JWTSecretKey string
var TTLToken int

// Cargar las variables de entorno y devolver configuración
func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error cargando el archivo .env: %v", err)
	}

	cfg := Config{
		MongoUser:     os.Getenv("MONGO_USER"),
		MongoPassword: os.Getenv("MONGO_PASSWORD"),
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoPort:     os.Getenv("MONGO_PORT"),
		MongoDB:       os.Getenv("MONGO_DB"),
		JWTSecretKey:  os.Getenv("JWT_SECRET_KEY"),
	}

	return cfg, nil
}

func Init() {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Leer la clave secreta desde las variables de entorno
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	if JWTSecretKey == "" {
		log.Fatal("JWT_SECRET_KEY no está configurado")
	}

	ttlTokenStr := os.Getenv("TTL_TOKEN")
	TTLToken, err = strconv.Atoi(ttlTokenStr)
	if err != nil {
		log.Fatalf("Error al convertir TTL_TOKEN a entero: %v", err)
	}
}

func SetupLogging() {
	// Crear o abrir el archivo de log
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Establecer el log en el archivo
	log.SetOutput(file)
	// Establecer un formato para los logs (opcional)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
