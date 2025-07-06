package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
	"time"
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
			VoteDate:   v.VoteDate.Format("2006-01-02"),
			ListNumber: v.ListNumber,
		})
	}
	return result, nil
}

func (u *listVoteUseCase) Add(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error) {
	date, _ := time.Parse("2006-01-02", dto.VoteDate)
	model := models.ListVote{
		VoteDate:   date,
		ListNumber: dto.ListNumber,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *listVoteUseCase) Update(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error) {
	date, _ := time.Parse("2006-01-02", dto.VoteDate)
	model := models.ListVote{
		ID:         dto.ID,
		VoteDate:   date,
		ListNumber: dto.ListNumber,
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
