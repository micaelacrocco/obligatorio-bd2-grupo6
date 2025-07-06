package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TableHandler struct {
	u interfaces.TableUseCase
}

func NewTableHandler(u interfaces.TableUseCase) *TableHandler {
	return &TableHandler{u: u}
}

func (h *TableHandler) GetAll(c *gin.Context) {
	tables, err := h.u.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch tables"})
		return
	}
	c.JSON(http.StatusOK, tables)
}

func (h *TableHandler) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	table, err := h.u.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "table not found"})
		return
	}
	c.JSON(http.StatusOK, table)
}

func (h *TableHandler) Add(c *gin.Context) {
	var dto dtos.TableDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	created, err := h.u.Add(dto)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not add table"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *TableHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.u.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "table not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
