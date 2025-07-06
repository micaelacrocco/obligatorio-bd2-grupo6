package interfaces

import "EleccionesUcu/models"

type ZoneRepository interface {
	GetAll() ([]models.Zone, error)
	GetById(id int) (*models.Zone, error)
	Add(zone models.Zone) (*models.Zone, error)
	Delete(id int) error
}
