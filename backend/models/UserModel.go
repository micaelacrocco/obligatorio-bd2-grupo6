package models

type User struct {
	ID             int
	PasswordHashed string
	UserType       string
	CitizenID      int
}
