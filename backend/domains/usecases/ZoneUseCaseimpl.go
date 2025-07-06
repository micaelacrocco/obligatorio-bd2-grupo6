package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type zoneUseCase struct {
	r interfaces.ZoneRepository
}

func NewZoneUseCase(r interfaces.ZoneRepository) interfaces.ZoneUseCase {
	return &zoneUseCase{r: r}
}

func (u *zoneUseCase) GetAll() ([]dtos.ZoneDto, error) {
	zones, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.ZoneDto
	for _, z := range zones {
		result = append(result, dtos.ZoneDto{
			ID:           z.ID,
			Name:         z.Name,
			Address:      z.Address,
			DepartmentID: z.DepartmentID,
		})
	}
	return result, nil
}

func (u *zoneUseCase) GetById(id int) (*dtos.ZoneDto, error) {
	zone, err := u.r.GetById(id)
	if err != nil {
		return nil, err
	}
	result := dtos.ZoneDto{
		ID:           zone.ID,
		Name:         zone.Name,
		Address:      zone.Address,
		DepartmentID: zone.DepartmentID,
	}
	return &result, nil
}

func (u *zoneUseCase) Add(dto dtos.ZoneDto) (*dtos.ZoneDto, error) {
	model := models.Zone{
		Name:         dto.Name,
		Address:      dto.Address,
		DepartmentID: dto.DepartmentID,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *zoneUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
