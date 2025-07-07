package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
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

func (c *circuitsUseCase) GetVotesPersonById(citizenID int) (*dtos.PersonVoteDTO, error) {
	personVote, err := c.r.GetVotesPersonById(citizenID)

	if err != nil {
		return nil, err
	}

	personVoteDTO := dtos.PersonVoteDTO{
		ID:         personVote.ID,
		VoteDate:   personVote.VoteDate,
		IsObserved: personVote.IsObserved,
		VoteType:   personVote.VoteType,
		CitizenID:  personVote.CitizenID,
		CircuitID:  personVote.CircuitID,
	}
	return &personVoteDTO, nil
}

func (c *circuitsUseCase) GetVotesByParty(circuitID int) ([]dtos.PartyVoteDto, error) {
	votes, err := c.r.GetVotesByParty(circuitID)
	if err != nil {
		return nil, err
	}

	totalVotes := 0
	for _, v := range votes {
		totalVotes += v.VoteCount
	}

	for i := range votes {
		percentage := float64(votes[i].VoteCount) / float64(totalVotes) * 100
		votes[i].VotePercentage = percentage
	}

	var voteDto []dtos.PartyVoteDto
	for _, v := range votes {
		voteDto = append(voteDto, dtos.PartyVoteDto{
			PartyName:  v.PartyName,
			Votes:      v.VoteCount,
			Percentage: v.VotePercentage,
		})
	}
	return voteDto, nil
}

func (c *circuitsUseCase) GetVotes(circuitID int) ([]dtos.CircuitResultDto, error) {
	votes, err := c.r.GetVotes(circuitID)
	if err != nil {
		return nil, err
	}

	totalVotes := 0
	for _, v := range votes {
		totalVotes += v.VoteCount
	}

	var voteDto []dtos.CircuitResultDto
	for _, v := range votes {
		percentage := 0.0
		if totalVotes > 0 {
			percentage = float64(v.VoteCount) / float64(totalVotes) * 100
		}

		listName := v.List
		if v.List == "" {
			// Si no hay número de lista, es un voto especial
			listName = v.PartyName // Puede ser "En Blanco", "Anulado", etc.
		}

		voteDto = append(voteDto, dtos.CircuitResultDto{
			List:       listName,
			PartyName:  v.PartyName,
			VoteCount:  v.VoteCount,
			Percentage: percentage,
		})
	}

	return voteDto, nil
}

func (c *circuitsUseCase) GetVotesByAllCandidates(circuitID int) ([]dtos.CircuitResultByAllCandidates, error) {
	votes, err := c.r.GetVotesByAllCandidates(circuitID)
	if err != nil {
		return nil, err
	}

	totalVotes := 0
	for _, v := range votes {
		totalVotes += v.VoteCount
	}

	var voteDto []dtos.CircuitResultByAllCandidates
	for _, v := range votes {
		percentage := 0.0
		if totalVotes > 0 {
			percentage = float64(v.VoteCount) / float64(totalVotes) * 100
		}

		partyName := v.Party
		if v.Party == "" {
			// Si no hay número de lista, es un voto especial
			partyName = v.Party // Puede ser "En Blanco", "Anulado", etc.
		}

		voteDto = append(voteDto, dtos.CircuitResultByAllCandidates{
			Party:      partyName,
			Candidate:  v.Candidate,
			VoteCount:  v.VoteCount,
			Percentage: percentage,
		})
	}

	return voteDto, nil
}

func (c *circuitsUseCase) GetMyCircuitByCitizenId(citizenId int) (*dtos.CircuitDto, error) {
	circuit, err := c.r.GetCircuitByCitizenId(citizenId)
	if err != nil {
		return nil, err
	}

	return &dtos.CircuitDto{
		ID:              circuit.ID,
		Location:        circuit.Location,
		Accessible:      circuit.Accessible,
		CredentialStart: circuit.CredentialStart,
		CredentialEnd:   circuit.CredentialEnd,
		PollingPlaceId:  circuit.PollingPlaceId,
	}, nil
}

func (c *circuitsUseCase) AddCircuit(circuit dtos.CircuitDto) (*dtos.CircuitDto, error) {
	circuitResult, err := c.r.AddCircuit(models.Circuit(circuit))
	if err != nil {
		return nil, err
	}

	circuitDto := dtos.CircuitDto{
		ID:              circuitResult.ID,
		Location:        circuitResult.Location,
		Accessible:      circuitResult.Accessible,
		CredentialStart: circuitResult.CredentialStart,
		CredentialEnd:   circuitResult.CredentialEnd,
		PollingPlaceId:  circuitResult.PollingPlaceId,
	}
	return &circuitDto, nil
}

func (c *circuitsUseCase) AddVotePerson(vote dtos.PersonVoteDTO) (*dtos.PersonVoteDTO, error) {
	personVoteResult, err := c.r.AddVotePerson(models.PersonVoteModel(vote))

	if err != nil {
		return nil, err
	}

	personVoteDto := dtos.PersonVoteDTO{
		ID:        personVoteResult.ID,
		VoteDate:  personVoteResult.VoteDate,
		VoteType:  personVoteResult.VoteType,
		CitizenID: personVoteResult.CitizenID,
		CircuitID: personVoteResult.CircuitID,
	}

	return &personVoteDto, nil
}
func (c *circuitsUseCase) Update(dto dtos.CircuitDto) (*dtos.CircuitDto, error) {
	updated, err := c.r.Update(models.Circuit(dto))
	if err != nil {
		return nil, err
	}
	result := dtos.CircuitDto(*updated)
	return &result, nil
}

func (c *circuitsUseCase) Delete(id int) error {
	return c.r.Delete(id)
}
