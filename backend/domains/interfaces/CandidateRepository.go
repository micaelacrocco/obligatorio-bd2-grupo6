package interfaces

import "EleccionesUcu/models"

type CandidateRepository interface {
	GetAll() ([]models.Candidate, error)
	GetByCitizenID(id int) ([]models.Candidate, error)
	Add(candidate models.Candidate) (*models.Candidate, error)
	Delete(citizenID int, listNumber int) error
}
