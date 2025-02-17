package handlers

import (
	"net/http"
	"notas_service/services"

	"github.com/gin-gonic/gin"
)

// NotaHandler estructura para manejar el servicio
type NotaHandler struct {
	notaService *services.NotaService
}

func NuevoNotaHandler(service *services.NotaService) *NotaHandler {
	return &NotaHandler{notaService: service}
}

func (h *NotaHandler) ObtenerNotas(c *gin.Context) {
	notas, err := h.notaService.ObtenerNotas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las notas"})
		return
	}
	c.JSON(http.StatusOK, notas)
}
