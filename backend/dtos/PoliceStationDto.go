package dtos

type PoliceStationDto struct {
	ID            int    `json:"id"`
	StationNumber int    `json:"station_number"`
	Address       string `json:"address"`
	DepartmentID  int    `json:"department_id"`
}
