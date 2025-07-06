package handlers

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/dtos"
	"EleccionesUcu/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type TableMemberHandler struct {
	u interfaces.TableMembersUseCase
}

func NewTableMemberHandler(u interfaces.TableMembersUseCase) *TableMemberHandler {
	return &TableMemberHandler{u: u}
}

func (h *TableMemberHandler) GetAll(c *gin.Context) {
	tableMembers, err := h.u.GetAll()

	if len(tableMembers) == 0 {
		log.Printf("error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "there are no table members"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch tables"})
		return
	}

	c.JSON(http.StatusOK, tableMembers)
	return
}

func (h *TableMemberHandler) GetCitizenIsTableMember(c *gin.Context) {

	citizenIdParam := c.Param("citizen_id")
	tableIdParam := c.Param("table_id")

	citizenId, citizenIdError := strconv.Atoi(citizenIdParam)
	tableId, tableIdError := strconv.Atoi(tableIdParam)

	if citizenIdError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the citizen id must be an integer"})
		return
	}
	if tableIdError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the table id must be an integer"})
		return
	}
	result, err := h.u.GetCitizenIsTableMember(citizenId, tableId)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func (h *TableMemberHandler) Add(c *gin.Context) {
	var tableMemberDto dtos.TableMembersDto

	if err := c.ShouldBindJSON(&tableMemberDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	tablesMembersResponse, err := h.u.Add(tableMemberDto)

	if errors.Is(err, utils.ErrForeignKeyNotFound) {
		c.JSON(http.StatusBadGateway, gin.H{"error": "not found FK"})
		return
	} else if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "There is a tableMember with the same citizen and table id"})
		return
	}

	c.JSON(http.StatusOK, tablesMembersResponse)
	return
}

func (h *TableMemberHandler) Delete(c *gin.Context) {

	citizenIdParam := c.Param("citizen_id")
	tableIdParam := c.Param("table_id")

	citizenId, citizenIdError := strconv.Atoi(citizenIdParam)
	tableId, tableIdError := strconv.Atoi(tableIdParam)

	if citizenIdError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the citizen id must be an integer"})
		return
	}
	if tableIdError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the table id must be an integer"})
		return
	}

	err := h.u.Delete(citizenId, tableId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "table member not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
