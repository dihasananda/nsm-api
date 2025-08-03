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

func (h *SchoolHandler) GetSchoolByNSM(c *gin.Context) {
	nsm := c.Query("nsm")
	if nsm == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Missing required query parameter 'nsm'",
		})
		return
	}

	// Validate NSM: must be exactly 10 digits
	if len(nsm) != 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "NSM must be exactly 10 digits",
		})
		return
	}
	for _, ch := range nsm {
		if ch < '0' || ch > '9' {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "NSM must contain only digits",
			})
			return
		}
	}

	// Query service
	school, err := h.service.GetSchoolByNSM(nsm)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "School with given NSM not found",
		})
		return
	}

	c.JSON(http.StatusOK, school)
}
