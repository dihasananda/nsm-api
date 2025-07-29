package handlers

import (
	"net/http"
	"nsm-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type NsmHandler struct {
	Repo *repository.NSMRepository
}

func NewNsmHandler(repo *repository.NSMRepository) *NsmHandler {
	return &NsmHandler{Repo: repo}
}

func (h *NsmHandler) GetNSMs(c *gin.Context) {
	nsms, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nsms)
}
