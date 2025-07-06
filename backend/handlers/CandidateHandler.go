package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CandidateHandler struct {
	u interfaces.CandidateUseCase
}

func NewCandidateHandler(u interfaces.CandidateUseCase) *CandidateHandler {
	return &CandidateHandler{u: u}
}

func (h *CandidateHandler) GetAll(c *gin.Context) {
	candidates, err := h.u.GetAll()

	if len(candidates) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no candidates"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get candidates"})
		return
	}
	c.JSON(http.StatusOK, candidates)
}

func (h *CandidateHandler) GetByCitizenID(c *gin.Context) {
	idParam := c.Param("citizen_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	result, err := h.u.GetByCitizenID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *CandidateHandler) Add(c *gin.Context) {
	var dto dtos.CandidateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add candidate"})
		log.Printf("error: %v", err)
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *CandidateHandler) Delete(c *gin.Context) {
	citizenID := c.Param("citizen_id")
	listNumber := c.Param("list_number")
	citizenIdInt, errCitizenIdParse := strconv.Atoi(citizenID)
	listNumberInt, errListNumberParse := strconv.Atoi(listNumber)

	if errCitizenIdParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Citizen Id must be an integer"})
		return
	}

	if errListNumberParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The List Number id must be an integer"})
		return
	}
	err := h.u.Delete(citizenIdInt, listNumberInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "candidate not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
