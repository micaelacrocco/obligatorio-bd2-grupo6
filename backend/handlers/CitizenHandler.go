package handlers

import (
	"EleccionesUcu/dtos"
	"EleccionesUcu/domains/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CitizenHandler struct {
	u interfaces.CitizenUseCase
}

func NewCitizenHandler(u interfaces.CitizenUseCase) *CitizenHandler {
	return &CitizenHandler{u: u}
}

func (h *CitizenHandler) GetAll(c *gin.Context) {
	citizens, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no citizens"})
		return
	}
	c.JSON(http.StatusOK, citizens)
}

func (h *CitizenHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the id must be an integer"})
		return
	}
	citizen, err := h.u.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no citizens with this id"})
		return
	}
	c.JSON(http.StatusOK, citizen)
}

func (h *CitizenHandler) AddCitizen(c *gin.Context) {
	var citizenDto dtos.CitizenDto

	if err := c.ShouldBindJSON(&citizenDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	citizenResponse, err := h.u.AddCitizen(citizenDto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "There is already a citizen with this id"})
		return
	}
	c.JSON(http.StatusCreated, citizenResponse)
}

func (h *CitizenHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the id must be an integer"})
		return
	}

	var citizenDto dtos.CitizenDto
	if err := c.ShouldBindJSON(&citizenDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = h.u.Update(id, citizenDto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no citizen with this id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Citizen updated successfully"})
}

func (h *CitizenHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the id must be an integer"})
		return
	}

	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There is no citizen with this id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Citizen deleted successfully"})
}
