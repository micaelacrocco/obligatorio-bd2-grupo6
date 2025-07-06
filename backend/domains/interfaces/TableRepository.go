package interfaces

import "EleccionesUcu/models"

type TableRepository interface {
	GetAll() ([]models.Table, error)
	GetById(id int) (*models.Table, error)
	Add(table models.Table) (*models.Table, error)
	Delete(id int) error
}
