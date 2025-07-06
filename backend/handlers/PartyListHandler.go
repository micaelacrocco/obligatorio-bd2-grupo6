package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PartyListHandler struct {
	u interfaces.PartyListUseCase
}

func NewPartyListHandler(u interfaces.PartyListUseCase) *PartyListHandler {
	return &PartyListHandler{u: u}
}

func (h *PartyListHandler) GetAll(c *gin.Context) {
	lists, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve party lists"})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *PartyListHandler) Add(c *gin.Context) {
	var dto dtos.PartyListDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add party list"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *PartyListHandler) Update(c *gin.Context) {
	var dto dtos.PartyListDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	updated, err := h.u.Update(dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "party list not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *PartyListHandler) Delete(c *gin.Context) {
	param := c.Param("list_number")
	listNumber, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid list_number"})
		return
	}
	err = h.u.Delete(listNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "party list not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
