package models

type User struct {
    ID           int
    CitizenID    int
    Username     string
    PasswordHash string
    UserType     string
}
