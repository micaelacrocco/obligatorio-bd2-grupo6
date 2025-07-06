package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ListVoteHandler struct {
	u interfaces.ListVoteUseCase
}

func NewListVoteHandler(u interfaces.ListVoteUseCase) *ListVoteHandler {
	return &ListVoteHandler{u: u}
}

func (h *ListVoteHandler) GetAll(c *gin.Context) {
	votes, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve list votes"})
		log.Printf("error, %v", err)
		return
	}
	c.JSON(http.StatusOK, votes)
}

func (h *ListVoteHandler) Add(c *gin.Context) {
	var dto dtos.ListVoteDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add list vote"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *ListVoteHandler) Update(c *gin.Context) {
	var dto dtos.ListVoteDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	updated, err := h.u.Update(dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list vote not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ListVoteHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "list vote not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
