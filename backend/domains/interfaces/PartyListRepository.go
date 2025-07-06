package interfaces

import "EleccionesUcu/models"

type PartyListRepository interface {
	GetAll() ([]models.PartyListModel, error)
	Add(list models.PartyListModel) (*models.PartyListModel, error)
	Update(list models.PartyListModel) (*models.PartyListModel, error)
	Delete(listNumber int) error
}
