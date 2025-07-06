package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type PoliceAgentHandler struct {
	u interfaces.PoliceAgentUseCase
}

func NewPoliceAgentHandler(u interfaces.PoliceAgentUseCase) *PoliceAgentHandler {
	return &PoliceAgentHandler{u: u}
}

func (h *PoliceAgentHandler) GetAll(c *gin.Context) {
	agents, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch agents"})
		return
	}
	c.JSON(http.StatusOK, agents)
}

func (h *PoliceAgentHandler) GetByCitizenID(c *gin.Context) {
	ciParam := c.Param("id")
	ci, err := strconv.Atoi(ciParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid CI"})
		return
	}

	agent, err := h.u.GetByCitizenID(ci)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "agent not found"})
		return
	}
	c.JSON(http.StatusOK, agent)
}

func (h *PoliceAgentHandler) Add(c *gin.Context) {
	var dto dtos.PoliceAgentDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add agent"})
		log.Printf("error %v", err)
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *PoliceAgentHandler) Update(c *gin.Context) {
	var dto dtos.PoliceAgentDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	updated, err := h.u.Update(dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not update agent"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *PoliceAgentHandler) Delete(c *gin.Context) {
	ciParam := c.Param("id")
	ci, err := strconv.Atoi(ciParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid CI"})
		return
	}
	err = h.u.Delete(ci)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "agent not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
