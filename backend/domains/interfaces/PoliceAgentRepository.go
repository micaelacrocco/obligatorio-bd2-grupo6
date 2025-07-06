package interfaces

import "EleccionesUcu/models"

type PoliceAgentRepository interface {
	GetAll() ([]models.PoliceAgent, error)
	GetByCitizenID(ci int) (*models.PoliceAgent, error)
	Add(agent models.PoliceAgent) (*models.PoliceAgent, error)
	Update(agent models.PoliceAgent) (*models.PoliceAgent, error)
	Delete(ci int) error
}
