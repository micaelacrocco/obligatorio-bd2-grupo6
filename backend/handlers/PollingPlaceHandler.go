package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PollingPlaceHandler struct {
	u interfaces.PollingPlaceUseCase
}

func NewPollingPlaceHandler(u interfaces.PollingPlaceUseCase) *PollingPlaceHandler {
	return &PollingPlaceHandler{u: u}
}

func (h *PollingPlaceHandler) GetAll(c *gin.Context) {
	all, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get polling places"})
		return
	}
	c.JSON(http.StatusOK, all)
}

func (h *PollingPlaceHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	result, err := h.u.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *PollingPlaceHandler) Add(c *gin.Context) {
	var dto dtos.PollingPlaceDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not create"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *PollingPlaceHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
