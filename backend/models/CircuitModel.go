package models

type Circuit struct {
	ID              int    `json:"id"`
	Location        string `json:"location"`
	Accesible       bool   `json:"is_accessible"`
	CredentialStart int    `json:"credential_start"`
	CredentialEnd   int    `json:"credential_end"`
	PollingPlaceId  int    `json:"polling_place_id"`
}
