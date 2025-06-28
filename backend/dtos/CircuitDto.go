package dtos

type CircuitDto struct {
	ID              int    `json:"id"`
	Location        string `json:"location"`
	Accessible      bool   `json:"is_accessible"`
	CredentialStart int    `json:"credential_start"`
	CredentialEnd   int    `json:"credential_end"`
	PollingPlaceId  int    `json:"polling_place_id"`
}
