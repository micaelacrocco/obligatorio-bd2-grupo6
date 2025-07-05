package dtos

type CitizenDto struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"` // Use string for JSON serialization
	Credential string `json:"credential"`
}
