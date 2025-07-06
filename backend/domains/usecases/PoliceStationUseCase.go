package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type policeStationUseCase struct {
	r interfaces.PoliceStationRepository
}

func NewPoliceStationUseCase(r interfaces.PoliceStationRepository) interfaces.PoliceStationUseCase {
	return &policeStationUseCase{r: r}
}

func (u *policeStationUseCase) GetAll() ([]dtos.PoliceStationDto, error) {
	stations, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.PoliceStationDto
	for _, s := range stations {
		result = append(result, dtos.PoliceStationDto{
			ID:            s.ID,
			StationNumber: s.StationNumber,
			Address:       s.Address,
			DepartmentID:  s.DepartmentID,
		})
	}
	return result, nil
}

func (u *policeStationUseCase) Add(dto dtos.PoliceStationDto) (*dtos.PoliceStationDto, error) {
	model := models.PoliceStation{
		StationNumber: dto.StationNumber,
		Address:       dto.Address,
		DepartmentID:  dto.DepartmentID,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *policeStationUseCase) Update(dto dtos.PoliceStationDto) (*dtos.PoliceStationDto, error) {
	model := models.PoliceStation{
		ID:            dto.ID,
		StationNumber: dto.StationNumber,
		Address:       dto.Address,
		DepartmentID:  dto.DepartmentID,
	}
	_, err := u.r.Update(model)
	if err != nil {
		return nil, err
	}
	return &dto, nil
}

func (u *policeStationUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
