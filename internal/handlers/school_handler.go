package handlers

import (
	"net/http"
	"nsm-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	Repo *repository.SchoolRepository
}

func NewSchoolHandler(repo *repository.SchoolRepository) *SchoolHandler {
	return &SchoolHandler{Repo: repo}
}

func (h *SchoolHandler) GetSchools(c *gin.Context) {
	schools, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}
