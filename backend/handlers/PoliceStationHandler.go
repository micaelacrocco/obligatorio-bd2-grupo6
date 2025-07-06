package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PoliceStationHandler struct {
	u interfaces.PoliceStationUseCase
}

func NewPoliceStationHandler(u interfaces.PoliceStationUseCase) *PoliceStationHandler {
	return &PoliceStationHandler{u: u}
}

func (h *PoliceStationHandler) GetAll(c *gin.Context) {
	stations, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch police stations"})
		return
	}
	c.JSON(http.StatusOK, stations)
}

func (h *PoliceStationHandler) Add(c *gin.Context) {
	var dto dtos.PoliceStationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add police station"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *PoliceStationHandler) Update(c *gin.Context) {
	var dto dtos.PoliceStationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	updated, err := h.u.Update(dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not update police station"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *PoliceStationHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "police station not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
