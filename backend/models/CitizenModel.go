package models

import "time"

type Citizen struct {
    ID           int
    FirstName    string
    LastName     string
    BirthDate    time.Time
    Credential   string
}
