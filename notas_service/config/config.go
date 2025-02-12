package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"strconv"
)

// Config estructura de configuración
type Config struct {
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPassword string
	MongoDB       string
}

var JWTSecretKey string
var TTLToken int

// Cargar las variables de entorno y devolver configuración
func LoadConfig() (Config, error) {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	cfg := Config{
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoPort:     os.Getenv("MONGO_PORT"),
		MongoUser:     os.Getenv("MONGO_USER"),
		MongoPassword: os.Getenv("MONGO_PASSWORD"),
		MongoDB:       os.Getenv("MONGO_DB"),
	}

	if cfg.MongoHost == "" || cfg.MongoPort == "" || cfg.MongoDB == "" {
		log.Fatalf("Las variables de entorno no están configuradas correctamente")
		return cfg, fmt.Errorf("missing required environment variables")
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
