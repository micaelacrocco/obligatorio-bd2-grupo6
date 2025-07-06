package interfaces

import "EleccionesUcu/models"

type DepartmentRepository interface {
	GetAll() ([]models.Department, error)
	Add(dept models.Department) (*models.Department, error)
	Delete(id int) error
}
