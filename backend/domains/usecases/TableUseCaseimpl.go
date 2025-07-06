package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type tableUseCase struct {
	r interfaces.TableRepository
}

func NewTableUseCase(r interfaces.TableRepository) interfaces.TableUseCase {
	return &tableUseCase{r: r}
}

func (u *tableUseCase) GetAll() ([]dtos.TableDto, error) {
	tables, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dtos.TableDto
	for _, t := range tables {
		result = append(result, dtos.TableDto{
			ID:        t.ID,
			CircuitID: t.CircuitID,
		})
	}
	return result, nil
}

func (u *tableUseCase) GetById(id int) (*dtos.TableDto, error) {
	table, err := u.r.GetById(id)
	if err != nil {
		return nil, err
	}
	dto := dtos.TableDto{
		ID:        table.ID,
		CircuitID: table.CircuitID,
	}
	return &dto, nil
}

func (u *tableUseCase) Add(dto dtos.TableDto) (*dtos.TableDto, error) {
	model := models.Table{CircuitID: dto.CircuitID}
	added, err := u.r.Add(model)
	if err != nil {
		return nil, err
	}
	dto.ID = added.ID
	return &dto, nil
}

func (u *tableUseCase) Delete(id int) error {
	return u.r.Delete(id)
}
