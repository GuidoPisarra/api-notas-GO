package controllers

import (
	"auth_service/config"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"auth_service/database"
	"auth_service/models"

	"log"
)

var jwtKey = []byte(config.JWTSecretKey) // Clave secreta para JWT

// Registro godoc
// @Summary Registra un nuevo usuario
// @Description Crea un nuevo usuario en el sistema
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param usuario body models.Usuario true "Usuario a registrar"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/registro [post]
func Registro(c *gin.Context) {
	var usuario models.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		log.Println("Error al vincular datos: ", err) // Log de error de validación
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Hashear la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al generar hash de contraseña: ", err) // Log de error al generar hash
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar hash"})
		return
	}
	usuario.Password = string(hash)

	// Guardar en la BD
	collection := database.GetCollection("usuarios")
	/* if err != nil {
		log.Println("Error al obtener la colección: ", err) // Log de error al obtener la colección
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la colección"})
		return
	} */
	_, err = collection.InsertOne(context.Background(), usuario)
	if err != nil {
		log.Println("Error al registrar usuario: ", err) // Log de error al guardar usuario
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
}

// Login godoc
// @Summary Inicia sesión de usuario
// @Description Verifica las credenciales del usuario y devuelve token
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param usuario body models.Usuario true "Datos de inicio de sesión"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var usuario models.Usuario
	var usuarioEncontrado models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		log.Printf("Error al vincular datos: %v\n", err) // Log de error de validación
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar usuario en la BD
	collection := database.GetCollection("usuarios")
	/* if err != nil {
		log.Printf("Error al obtener la colección: %v\n", err) // Log de error al obtener la colección
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la colección"})
		return
	} */
	err := collection.FindOne(context.Background(), bson.M{"email": usuario.Email}).Decode(&usuarioEncontrado)
	if err != nil {
		log.Printf("Error al encontrar usuario: %v\n", err) // Log de error al encontrar usuario
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Comparar contraseñas
	err = bcrypt.CompareHashAndPassword([]byte(usuarioEncontrado.Password), []byte(usuario.Password))
	if err != nil {
		log.Printf("Error al comparar contraseñas: %v\n", err) // Log de error al comparar contraseñas
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Crear JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": usuarioEncontrado.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(config.TTLToken)).Unix(), // Expira en 24 horas
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error al generar el token: %v\n", err) // Log de error al generar el token
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
