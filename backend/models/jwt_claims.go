package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	CitizenID  int    `json:"citizen_id"`
	Credential string `json:"credential"`
	UserType   string `json:"user_type"`
	jwt.RegisteredClaims
}
