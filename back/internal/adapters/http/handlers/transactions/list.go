package transactions

import (
	"magnifin/internal/adapters"
	"magnifin/internal/adapters/http/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) List(c *gin.Context) {
	user := middlewares.GetUser(c.Request.Context())
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	fromStr := c.Query("from")
	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
		return
	}

	toStr := c.Query("to")
	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
		return
	}

	transactions, err := h.Service.GetAllByUserBetweenDates(c.Request.Context(), user, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t := make([]adapters.EnrichedTransaction, len(transactions))
	for i, tx := range transactions {
		t[i] = *h.Mapper.ToPublicFormat(&tx)
	}

	c.JSON(http.StatusOK, listResponse{Transactions: t})
}

type listResponse struct {
	Transactions []adapters.EnrichedTransaction `json:"transactions"`
}
