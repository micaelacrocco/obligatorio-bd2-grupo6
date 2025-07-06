package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type pollingPlaceUseCase struct {
	r interfaces.PollingPlaceRepository
}

func NewPollingPlaceUseCase(r interfaces.PollingPlaceRepository) interfaces.PollingPlaceUseCase {
	return &pollingPlaceUseCase{r: r}
}

func (u *pollingPlaceUseCase) GetAll() ([]dtos.PollingPlaceDto, error) {
	all, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.PollingPlaceDto
	for _, p := range all {
		result = append(result, dtos.PollingPlaceDto{
			ID:      p.ID,
			Name:    p.Name,
			Type:    p.Type,
			Address: p.Address,
			ZoneID:  p.ZoneID,
		})
	}
	return result, nil
}

func (u *pollingPlaceUseCase) GetByID(id int) (*dtos.PollingPlaceDto, error) {
	p, err := u.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dtos.PollingPlaceDto{
		ID:      p.ID,
		Name:    p.Name,
		Type:    p.Type,
		Address: p.Address,
		ZoneID:  p.ZoneID,
	}, nil
}

func (u *pollingPlaceUseCase) Add(dto dtos.PollingPlaceDto) (*dtos.PollingPlaceDto, error) {
	model := models.PollingPlace{
		Name:    dto.Name,
		Type:    dto.Type,
		Address: dto.Address,
		ZoneID:  dto.ZoneID,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *pollingPlaceUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
