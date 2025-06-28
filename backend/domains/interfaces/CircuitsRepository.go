package interfaces

import (
	"EleccionesUcu/models"
)

type CircuitsRepository interface {
	GetAll() ([]models.Circuit, error)
}
