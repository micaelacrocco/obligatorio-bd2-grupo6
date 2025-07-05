package interfaces

import (
	"EleccionesUcu/models"
)

type CitizenRepository interface {
	GetByID(id int) (models.Citizen, error)
	GetAll() ([]models.Citizen, error)
	Create(citizen models.Citizen) (int, error)
	Update(id int, citizen models.Citizen) error
	Delete(id int) error
}