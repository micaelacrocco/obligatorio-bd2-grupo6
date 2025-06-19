package dtos

type CircuitDto struct {
	ID              int
	Location        string
	Accesible       bool
	CredentialStart int
	CredentialEnd   int
	PollingPlaceId  int
}
