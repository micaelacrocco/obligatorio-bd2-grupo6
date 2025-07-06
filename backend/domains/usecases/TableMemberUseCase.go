package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/models"
)

type tableMemberUseCase struct {
	r interfaces.TableMembersRepository
}

func NewTableMemberUseCase(r interfaces.TableMembersRepository) interfaces.TableMembersUseCase {
	return &tableMemberUseCase{r: r}
}

func (u tableMemberUseCase) GetAll() ([]dtos.TableMembersDto, error) {
	tableMembers, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dtos.TableMembersDto
	for _, t := range tableMembers {
		result = append(result, dtos.TableMembersDto{
			TableID:         t.TableID,
			CitizenID:       t.CitizenID,
			IntegrationDate: t.IntegrationDate,
			Duty:            t.Duty,
		})
	}
	return result, nil
}

func (u tableMemberUseCase) GetCitizenIsTableMember(citizenID int, tableID int) (bool, error) {
	_, err := u.r.GetCitizenIsTableMember(tableID, citizenID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u tableMemberUseCase) Add(tableMember dtos.TableMembersDto) (*dtos.TableMembersDto, error) {
	result, err := u.r.Add(models.TableMembers(tableMember))
	if err != nil {
		return nil, err
	}

	return &dtos.TableMembersDto{
		TableID:         result.TableID,
		CitizenID:       result.CitizenID,
		IntegrationDate: result.IntegrationDate,
		Duty:            result.Duty,
	}, nil
}

func (u tableMemberUseCase) Delete(citizenID int, tableMemberID int) error {
	return u.r.Delete(citizenID, tableMemberID)
}
