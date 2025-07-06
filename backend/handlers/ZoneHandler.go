package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ZoneHandler struct {
	u interfaces.ZoneUseCase
}

func NewZoneHandler(u interfaces.ZoneUseCase) *ZoneHandler {
	return &ZoneHandler{u: u}
}

func (h *ZoneHandler) GetAll(c *gin.Context) {
	zones, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch zones"})
		return
	}
	c.JSON(http.StatusOK, zones)
}

func (h *ZoneHandler) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	zone, err := h.u.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
		return
	}
	c.JSON(http.StatusOK, zone)
}

func (h *ZoneHandler) Add(c *gin.Context) {
	var dto dtos.ZoneDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add zone"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *ZoneHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
