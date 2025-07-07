package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	CitizenID  int    `json:"citizen_id"`
	Credential string `json:"credential"`
	CircuitId  int    `json:"circuit_id"`
	jwt.RegisteredClaims
}
