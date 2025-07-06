package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type departmentUseCase struct {
	r interfaces.DepartmentRepository
}

func NewDepartmentUseCase(r interfaces.DepartmentRepository) interfaces.DepartmentUseCase {
	return &departmentUseCase{r: r}
}

func (u *departmentUseCase) GetAll() ([]dtos.DepartmentDto, error) {
	depts, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.DepartmentDto
	for _, d := range depts {
		result = append(result, dtos.DepartmentDto{
			ID:   d.ID,
			Name: d.Name,
		})
	}
	return result, nil
}

func (u *departmentUseCase) Add(dto dtos.DepartmentDto) (*dtos.DepartmentDto, error) {
	dept := models.Department{Name: dto.Name}
	added, err := u.r.Add(dept)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *departmentUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
