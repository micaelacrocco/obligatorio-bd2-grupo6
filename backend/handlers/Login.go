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

/* ------------------------------------------------------------------------- */
/*  Estructura y constructor                                                 */
/* ------------------------------------------------------------------------- */

type LoginHandler struct {
	citizenUC interfaces.CitizenUseCase
	userUC    interfaces.UserUseCase
}

func NewLoginHandler(citizenUC interfaces.CitizenUseCase,
	userUC interfaces.UserUseCase) *LoginHandler {
	return &LoginHandler{citizenUC: citizenUC, userUC: userUC}
}

/* ------------------------------------------------------------------------- */
/*  DTO recibido desde el frontend                                           */
/* ------------------------------------------------------------------------- */

type loginRequest struct {
	CI         int    `json:"ci"`         // cédula sin puntos ni guiones
	Credential string `json:"credential"` // ABC12345
	Password   string `json:"password"`   // texto plano
}

/* ------------------------------------------------------------------------- */
/*  Endpoint POST /auth/login                                                */
/* ------------------------------------------------------------------------- */

func (h *LoginHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	/* ---------- 1 · Validar ciudadano (CI + credencial) ---------- */
	citizen, err := h.citizenUC.GetByID(req.CI)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ci_not_found"})
		return
	}
	if citizen.Credential != req.Credential {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credential_mismatch"})
		return
	}

	/* ---------- 2 · Validar usuario y contraseña ---------- */
	user, err := h.userUC.FindByCitizenID(citizen.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_not_found"})
		return
	}
	if !h.userUC.CheckPassword(req.Password, user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong_password"})
		return
	}

	/* ---------- 3 · Generar JWT ---------- */
	claims := models.JwtCustomClaims{
		CitizenID:  citizen.ID,
		Credential: citizen.Credential,
		UserType:   user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(configuration.JWT_SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token_sign_error"})
		return
	}

	/* ---------- 4 · Respuesta ---------- */
	c.JSON(http.StatusOK, gin.H{
		"token":     signed,
		"user_type": user.UserType,
		"ci":        citizen.ID,
	})
}
