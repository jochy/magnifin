package transactions

import (
	"magnifin/internal/adapters/http/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) MinMax(c *gin.Context) {
	user := middlewares.GetUser(c.Request.Context())
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	minMax, err := h.Service.GetTransactionMinMaxDateByUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if minMax == nil {
		c.JSON(http.StatusOK, minMaxResponse{})
		return
	}

	c.JSON(http.StatusOK, minMaxResponse{
		Min: &minMax.Min,
		Max: &minMax.Max,
	})
}

type minMaxResponse struct {
	Min *time.Time `json:"min"`
	Max *time.Time `json:"max"`
}
