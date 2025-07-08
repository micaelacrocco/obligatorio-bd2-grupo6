package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type listVoteUseCase struct {
	r interfaces.ListVoteRepository
}

func NewListVoteUseCase(r interfaces.ListVoteRepository) interfaces.ListVoteUseCase {
	return &listVoteUseCase{r: r}
}

func (u *listVoteUseCase) GetAll() ([]dtos.ListVoteDto, error) {
	votes, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.ListVoteDto
	for _, v := range votes {
		result = append(result, dtos.ListVoteDto{
			ID:         v.ID,
			VoteDate:   v.VoteDate,
			ListNumber: v.ListNumber,
			CircuitID:  v.CircuitID,
		})
	}
	return result, nil
}

func (u *listVoteUseCase) Add(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error) {
	model := models.ListVote{
		VoteDate:   dto.VoteDate,
		ListNumber: dto.ListNumber,
		CircuitID:  dto.CircuitID,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *listVoteUseCase) Update(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error) {
	model := models.ListVote{
		ID:         dto.ID,
		VoteDate:   dto.VoteDate,
		ListNumber: dto.ListNumber,
		CircuitID:  dto.CircuitID,
	}
	_, err := u.r.Update(model)
	if err != nil {
		return nil, err
	}
	return &dto, nil
}

func (u *listVoteUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
