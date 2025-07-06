package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type politicalPartyUseCase struct {
	r interfaces.PoliticalPartyRepository
}

func NewPoliticalPartyUseCase(r interfaces.PoliticalPartyRepository) interfaces.PoliticalPartyUseCase {
	return &politicalPartyUseCase{r: r}
}

func (u politicalPartyUseCase) GetAll() ([]dtos.PoliticalPartyDTO, error) {
	parties, err := u.r.GetAll()

	if err != nil {
		return nil, err
	}

	var partiesDto []dtos.PoliticalPartyDTO

	for _, p := range parties {
		partiesDto = append(partiesDto, dtos.PoliticalPartyDTO{
			ID:   p.ID,
			Name: p.Name,
		})
	}

	return partiesDto, nil
}

func (u politicalPartyUseCase) Add(party dtos.PoliticalPartyDTO) (*dtos.PoliticalPartyDTO, error) {
	partyResult, err := u.r.Add(models.PoliticalParty(party))

	if err != nil {
		return nil, err
	}

	partyDto := dtos.PoliticalPartyDTO{
		ID:   partyResult.ID,
		Name: partyResult.Name,
	}

	return &partyDto, nil
}

func (u politicalPartyUseCase) Update(party dtos.PoliticalPartyDTO) (*dtos.PoliticalPartyDTO, error) {
	updatedParty, err := u.r.Update(models.PoliticalParty(party))

	if err != nil {
		return nil, err
	}

	partyDto := dtos.PoliticalPartyDTO{
		ID:   updatedParty.ID,
		Name: updatedParty.Name,
	}
	return &partyDto, nil
}

func (u politicalPartyUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
