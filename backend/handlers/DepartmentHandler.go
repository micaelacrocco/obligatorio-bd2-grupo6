package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	u interfaces.DepartmentUseCase
}

func NewDepartmentHandler(u interfaces.DepartmentUseCase) *DepartmentHandler {
	return &DepartmentHandler{u: u}
}

func (h *DepartmentHandler) GetAll(c *gin.Context) {
	depts, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve departments"})
		return
	}
	c.JSON(http.StatusOK, depts)
}

func (h *DepartmentHandler) Add(c *gin.Context) {
	var dto dtos.DepartmentDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add department"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *DepartmentHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
