package dtos

type ZoneDto struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DepartmentID int    `json:"department_id"`
}
