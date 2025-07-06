package interfaces

import "EleccionesUcu/models"

type PoliceStationRepository interface {
	GetAll() ([]models.PoliceStation, error)
	Add(station models.PoliceStation) (*models.PoliceStation, error)
	Update(station models.PoliceStation) (*models.PoliceStation, error)
	Delete(id int) error
}
