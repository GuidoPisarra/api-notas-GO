package services

import (
	"notas_service/models"
	"notas_service/repository"
)

// NotaService estructura que usa el repositorio
type NotaService struct {
	notaRepo *repository.NotaRepository
}

// NuevoNotaService crea una nueva instancia del servicio
func NuevoNotaService(repo *repository.NotaRepository) *NotaService {
	return &NotaService{notaRepo: repo}
}

// ObtenerNotas obtiene todas las notas desde el repositorio
func (s *NotaService) ObtenerNotas() ([]models.Nota, error) {
	return s.notaRepo.ObtenerNotas()
}
