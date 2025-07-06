package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PoliticalPartyHandler struct {
	u interfaces.PoliticalPartyUseCase
}

func NewPoliticalPartyHandler(u interfaces.PoliticalPartyUseCase) *PoliticalPartyHandler {
	return &PoliticalPartyHandler{u: u}
}

func (h *PoliticalPartyHandler) GetAll(c *gin.Context) {
	parties, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve political parties"})
		return
	}
	c.JSON(http.StatusOK, parties)
}

func (h *PoliticalPartyHandler) Add(c *gin.Context) {
	var party dtos.PoliticalPartyDTO

	if err := c.ShouldBindJSON(&party); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	created, err := h.u.Add(party)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add political party"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *PoliticalPartyHandler) Update(c *gin.Context) {
	var party dtos.PoliticalPartyDTO

	if err := c.ShouldBindJSON(&party); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	updated, err := h.u.Update(party)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not update political party"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *PoliticalPartyHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "political party not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
