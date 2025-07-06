package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
	"time"
)

type candidateUseCase struct {
	r interfaces.CandidateRepository
}

func NewCandidateUseCase(r interfaces.CandidateRepository) interfaces.CandidateUseCase {
	return &candidateUseCase{r: r}
}

func (u *candidateUseCase) GetAll() ([]dtos.CandidateDto, error) {
	all, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.CandidateDto
	for _, c := range all {
		result = append(result, dtos.CandidateDto{
			CitizenID:     c.CitizenID,
			ListNumber:    c.ListNumber,
			StartDate:     c.StartDate.Format("2006-01-02"),
			EndDate:       c.EndDate.Format("2006-01-02"),
			CandidacyType: c.CandidacyType,
		})
	}
	return result, nil
}

func (u *candidateUseCase) GetByCitizenID(id int) ([]dtos.CandidateDto, error) {
	list, err := u.r.GetByCitizenID(id)
	if err != nil {
		return nil, err
	}
	var result []dtos.CandidateDto
	for _, c := range list {
		result = append(result, dtos.CandidateDto{
			CitizenID:     c.CitizenID,
			ListNumber:    c.ListNumber,
			StartDate:     c.StartDate.Format("2006-01-02"),
			EndDate:       c.EndDate.Format("2006-01-02"),
			CandidacyType: c.CandidacyType,
		})
	}
	return result, nil
}

func (u *candidateUseCase) Add(dto dtos.CandidateDto) (*dtos.CandidateDto, error) {
	start, _ := time.Parse("2006-01-02", dto.StartDate)
	end, _ := time.Parse("2006-01-02", dto.EndDate)

	model := models.Candidate{
		CitizenID:     dto.CitizenID,
		ListNumber:    dto.ListNumber,
		StartDate:     start,
		EndDate:       end,
		CandidacyType: dto.CandidacyType,
	}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.StartDate = added.StartDate.Format("2006-01-02")
	dto.EndDate = added.EndDate.Format("2006-01-02")
	return &dto, nil
}

func (u *candidateUseCase) Delete(citizenID int, listNumber int) error {
	return u.r.Delete(citizenID, listNumber)
}
