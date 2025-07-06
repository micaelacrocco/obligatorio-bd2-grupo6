package dtos

type TableMembersDto struct {
	TableID         int    `json:"table_id"`
	CitizenID       int    `json:"citizen_id"`
	IntegrationDate string `json:"integration_date"`
	Duty            string `json:"duty"`
}
