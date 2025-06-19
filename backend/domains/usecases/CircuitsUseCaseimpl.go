package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
)

type circuitsUseCase struct {
	r interfaces.CircuitsRepository
}

func NewCircuitsUseCase(r interfaces.CircuitsRepository) interfaces.CircuitsUseCase {
	return &circuitsUseCase{r: r}
}

func (c *circuitsUseCase) GetAll() ([]dtos.CircuitDto, error) {
	circuits, err := c.r.GetAll()

	if err != nil {
		return nil, err
	}
	var circuitsDto []dtos.CircuitDto

	for _, c := range circuits {
		circuitsDto = append(circuitsDto, dtos.CircuitDto{
			ID:              c.ID,
			Location:        c.Location,
			Accessible:      c.Accessible,
			CredentialStart: c.CredentialStart,
			CredentialEnd:   c.CredentialEnd,
			PollingPlaceId:  c.PollingPlaceId,
		})
	}
	return circuitsDto, nil
}

func (c *circuitsUseCase) GetById(id int) (*dtos.CircuitDto, error) {
	circuit, err := c.r.GetById(id)

	if err != nil {
		return nil, err
	}

	circuitDto := dtos.CircuitDto{
		ID:              circuit.ID,
		Location:        circuit.Location,
		Accessible:      circuit.Accessible,
		CredentialStart: circuit.CredentialStart,
		CredentialEnd:   circuit.CredentialEnd,
		PollingPlaceId:  circuit.PollingPlaceId,
	}
	return &circuitDto, nil
}
