package controllers

import (
	"context"
	"net/http"
	"time"

	"api-notas-Go/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"api-notas-Go/database"
	"api-notas-Go/models"
)

var jwtKey = []byte(config.JWTSecretKey) // Clave secreta para JWT

// Registro de usuario
func Registro(c *gin.Context) {
	var usuario models.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Hashear la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar hash"})
		return
	}
	usuario.Password = string(hash)

	// Guardar en la BD
	collection := database.GetCollection("usuarios")
	_, err = collection.InsertOne(context.Background(), usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
}

// Login y generación de JWT
func Login(c *gin.Context) {
	var usuario models.Usuario
	var usuarioEncontrado models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar usuario en la BD
	collection := database.GetCollection("usuarios")
	err := collection.FindOne(context.Background(), bson.M{"email": usuario.Email}).Decode(&usuarioEncontrado)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Comparar contraseñas
	err = bcrypt.CompareHashAndPassword([]byte(usuarioEncontrado.Password), []byte(usuario.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Crear JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": usuarioEncontrado.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
