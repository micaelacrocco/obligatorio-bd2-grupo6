package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type citizenUseCase struct {
	r interfaces.CitizenRepository
}

func NewCitizenUseCase(r interfaces.CitizenRepository) interfaces.CitizenUseCase {
	return &citizenUseCase{r: r}
}

func (c *citizenUseCase) GetAll() ([]dtos.CitizenDto, error) {
	citizens, err := c.r.GetAll()
	if err != nil {
		return nil, err
	}

	var citizensDto []dtos.CitizenDto
	for _, citizen := range citizens {
		citizensDto = append(citizensDto, dtos.CitizenDto{
			ID:         citizen.ID,
			FirstName:  citizen.FirstName,
			LastName:   citizen.LastName,
			BirthDate:  citizen.BirthDate,
			Credential: citizen.Credential,
		})
	}

	return citizensDto, nil
}

func (c *citizenUseCase) GetByID(id int) (dtos.CitizenDto, error) {
	citizen, err := c.r.GetByID(id)
	if err != nil {
		return dtos.CitizenDto{}, err
	}

	return dtos.CitizenDto{
		ID:         citizen.ID,
		FirstName:  citizen.FirstName,
		LastName:   citizen.LastName,
		BirthDate:  citizen.BirthDate,
		Credential: citizen.Credential,
	}, nil
}

func (c *citizenUseCase) AddCitizen(dto dtos.CitizenDto) (dtos.CitizenDto, error) {

	citizen := models.Citizen{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		BirthDate:  dto.BirthDate,
		Credential: dto.Credential,
	}

	id, err := c.r.Create(citizen)
	if err != nil {
		return dtos.CitizenDto{}, err
	}
	dto.ID = id
	return dto, nil
}

func (c *citizenUseCase) Update(id int, dto dtos.CitizenDto) error {

	citizen := models.Citizen{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		BirthDate:  dto.BirthDate,
		Credential: dto.Credential,
	}

	return c.r.Update(id, citizen)
}

func (c *citizenUseCase) Delete(id int) error {
	return c.r.Delete(id)
}
