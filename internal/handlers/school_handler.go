package handlers

import (
	"net/http"
	"nsm-api/internal/services"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	service *services.SchoolService
}

func NewSchoolHandler(service *services.SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

func (h *SchoolHandler) GetSchools(c *gin.Context) {
	schools, err := h.service.GetAllSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}
