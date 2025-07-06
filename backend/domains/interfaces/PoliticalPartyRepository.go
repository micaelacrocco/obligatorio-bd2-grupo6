package interfaces

import "EleccionesUcu/models"

type PoliticalPartyRepository interface {
	GetAll() ([]models.PoliticalParty, error)
	Add(party models.PoliticalParty) (*models.PoliticalParty, error)
	Update(party models.PoliticalParty) (*models.PoliticalParty, error)
	Delete(id int) error
}
