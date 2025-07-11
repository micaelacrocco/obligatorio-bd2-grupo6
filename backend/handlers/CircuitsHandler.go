package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
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

func (h *CircuitsHandler) GetVotesByParty(c *gin.Context) {
	circuitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid circuit ID"})
		return
	}

	votes, err := h.u.GetVotesByParty(circuitID)

	if len(votes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "this circuit doesnt have votes"})
		return
	}
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch vote data"})
		return
	}

	c.JSON(http.StatusOK, votes)
}

func (h *CircuitsHandler) GetVotes(c *gin.Context) {
	circuitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid circuit ID"})
		return
	}

	votes, err := h.u.GetVotes(circuitID)

	if len(votes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "this circuit doesnt have votes"})
		return
	}
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch vote data"})
		return
	}

	c.JSON(http.StatusOK, votes)
}

func (h *CircuitsHandler) GetVotesByAllCandidates(c *gin.Context) {
	circuitID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid circuit ID"})
		return
	}

	votes, err := h.u.GetVotesByAllCandidates(circuitID)
	log.Printf("error: %v", votes)
	if len(votes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "this circuit doesnt have votes"})
		return
	}
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch vote data"})
		return
	}

	c.JSON(http.StatusOK, votes)
}

func (h *CircuitsHandler) AddCircuit(c *gin.Context) {
	var circuit dtos.CircuitDto

	if err := c.ShouldBindJSON(&circuit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	circuitResponse, err := h.u.AddCircuit(circuit)
	if errors.Is(err, utils.ErrForeignKeyNotFound) {
		c.JSON(http.StatusBadGateway, gin.H{"error": "not found FK"})
		return
	} else if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "There is already a circuit with this id"})
		return
	}

	c.JSON(http.StatusOK, circuitResponse)
	return
}

func (h *CircuitsHandler) AddPersonVote(c *gin.Context) {
	var personVote dtos.PersonVoteDTO

	if err := c.ShouldBindJSON(&personVote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	person, _ := h.u.GetVotesPersonById(personVote.CitizenID)

	if person != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "This person already voted"})
		return
	}

	personVoteResponse, err := h.u.AddVotePerson(personVote)
	if errors.Is(err, utils.ErrForeignKeyNotFound) {
		c.JSON(http.StatusBadGateway, gin.H{"error": "not found FK"})
		return

	} else if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "There is already a vote with this id"})
		return
	}

	c.JSON(http.StatusOK, personVoteResponse)
	return
}

func (h *CircuitsHandler) Update(c *gin.Context) {
	var dto dtos.CircuitDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	updated, err := h.u.Update(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not update"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *CircuitsHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		log.Printf("error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
