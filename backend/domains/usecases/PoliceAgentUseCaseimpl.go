package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type policeAgentUseCase struct {
	r interfaces.PoliceAgentRepository
}

func NewPoliceAgentUseCase(r interfaces.PoliceAgentRepository) interfaces.PoliceAgentUseCase {
	return &policeAgentUseCase{r: r}
}

func (u *policeAgentUseCase) GetAll() ([]dtos.PoliceAgentDto, error) {
	agents, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.PoliceAgentDto
	for _, a := range agents {
		result = append(result, dtos.PoliceAgentDto{
			CitizenID:       a.CitizenID,
			PoliceStationID: a.PoliceStationID,
			PollingPlaceID:  a.PollingPlaceID,
		})
	}
	return result, nil
}

func (u *policeAgentUseCase) GetByCitizenID(ci int) (*dtos.PoliceAgentDto, error) {
	agent, err := u.r.GetByCitizenID(ci)
	if err != nil {
		return nil, err
	}
	dto := dtos.PoliceAgentDto{
		CitizenID:       agent.CitizenID,
		PoliceStationID: agent.PoliceStationID,
		PollingPlaceID:  agent.PollingPlaceID,
	}
	return &dto, nil
}

func (u *policeAgentUseCase) Add(dto dtos.PoliceAgentDto) (*dtos.PoliceAgentDto, error) {
	model := models.PoliceAgent(dto)
	created, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	result := dtos.PoliceAgentDto(*created)
	return &result, nil
}

func (u *policeAgentUseCase) Update(dto dtos.PoliceAgentDto) (*dtos.PoliceAgentDto, error) {
	model := models.PoliceAgent(dto)
	updated, err := u.r.Update(model)
	if err != nil {
		return nil, err
	}
	result := dtos.PoliceAgentDto(*updated)
	return &result, nil
}

func (u *policeAgentUseCase) Delete(ci int) error {
	return u.r.Delete(ci)
}
