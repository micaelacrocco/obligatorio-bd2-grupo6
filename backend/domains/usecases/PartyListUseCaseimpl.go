package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type partyListUseCase struct {
	r interfaces.PartyListRepository
}

func NewPartyListUseCase(r interfaces.PartyListRepository) interfaces.PartyListUseCase {
	return &partyListUseCase{r: r}
}

func (u *partyListUseCase) GetAll() ([]dtos.PartyListDto, error) {
	lists, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.PartyListDto
	for _, l := range lists {
		result = append(result, dtos.PartyListDto{
			ListNumber: l.ListNumber,
			PartyID:    l.PartyID,
		})
	}
	return result, nil
}

func (u *partyListUseCase) Add(dto dtos.PartyListDto) (*dtos.PartyListDto, error) {
	list, err := u.r.Add(models.PartyListModel(dto))
	if err != nil {
		return nil, err
	}
	result := dtos.PartyListDto(*list)
	return &result, nil
}

func (u *partyListUseCase) Update(dto dtos.PartyListDto) (*dtos.PartyListDto, error) {
	list, err := u.r.Update(models.PartyListModel(dto))
	if err != nil {
		return nil, err
	}
	result := dtos.PartyListDto(*list)
	return &result, nil
}

func (u *partyListUseCase) Delete(listNumber int) error {
	return u.r.Delete(listNumber)
}
