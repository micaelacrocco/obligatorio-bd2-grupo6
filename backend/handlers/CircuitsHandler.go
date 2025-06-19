package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *CircuitsHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the id must be an integer"})
		return
	}
	circuit, err := h.u.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no circuits with this id"})
		return
	}
	c.JSON(http.StatusOK, circuit)
	return
}
