package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CircuitsHandler struct {
	u interfaces.CircuitsUseCase
}

func NewCircuitsHandler(u interfaces.CircuitsUseCase) *CircuitsHandler {
	return &CircuitsHandler{u: u}
}

func (h *CircuitsHandler) GetAll(c *gin.Context) {
	circuits, err := h.u.GetAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no circuits"})
		return
	}
	c.JSON(http.StatusOK, circuits)
	return
}
