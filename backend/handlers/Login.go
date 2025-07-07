package handlers

import (
	"EleccionesUcu/configuration"
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type LoginHandler struct {
	u  interfaces.CitizenUseCase
	cu interfaces.CircuitsUseCase
}

func NewLoginHandler(u interfaces.CitizenUseCase, cu interfaces.CircuitsUseCase) *LoginHandler {
	return &LoginHandler{u: u, cu: cu}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var body struct {
		CitizenID int `json:"citizen_id"`
		CircuitID int `json:"circuit_id"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// Verifica si el ciudadano existe
	citizen, citizenErr := h.u.GetByID(body.CitizenID)
	//Verifica si el circuito existe
	circuit, circuitErr := h.cu.GetById(body.CircuitID)

	if citizenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Citizen not found"})
		return
	}
	if circuitErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Circuit not found"})
		return
	}

	// Genera el token si el ciudadano existe
	claims := models.JwtCustomClaims{
		CitizenID:  citizen.ID,
		Credential: citizen.Credential,
		CircuitId:  circuit.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(configuration.JWT_SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not sign token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}
