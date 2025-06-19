package interfaces

import (
	"EleccionesUcu/models"
)

type CircuitsRepository interface {
	GetAll() ([]models.Circuit, error)
	GetById(id int) (*models.Circuit, error)
}
