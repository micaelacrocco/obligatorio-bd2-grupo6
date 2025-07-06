package interfaces

import "EleccionesUcu/models"

type PollingPlaceRepository interface {
	GetAll() ([]models.PollingPlace, error)
	GetByID(id int) (*models.PollingPlace, error)
	Add(models.PollingPlace) (*models.PollingPlace, error)
	Delete(id int) error
}
