package dtos

type CircuitDto struct {
	ID              int
	Location        string
	Accessible      bool
	CredentialStart int
	CredentialEnd   int
	PollingPlaceId  int
}
