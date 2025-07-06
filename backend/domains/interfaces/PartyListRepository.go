package interfaces

import "EleccionesUcu/models"

type PartyListRepository interface {
	GetAll() ([]models.PartyList, error)
	Add(list models.PartyList) (*models.PartyList, error)
	Update(list models.PartyList) (*models.PartyList, error)
	Delete(listNumber int) error
}
