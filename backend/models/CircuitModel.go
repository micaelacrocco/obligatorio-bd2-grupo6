package models

type Circuit struct {
	ID              int
	Location        string
	Accessible      bool
	CredentialStart int
	CredentialEnd   int
	PollingPlaceId  int
}
