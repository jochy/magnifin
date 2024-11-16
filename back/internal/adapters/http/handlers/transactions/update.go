package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Update(c *gin.Context) {
	var req updateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.Service.Update(c.Request.Context(), req.ID, req.Category, req.UserCounterparty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	publicTrs := h.toPublicFormat(tx)
	c.JSON(http.StatusOK, publicTrs)
}

type updateTransactionRequest struct {
	ID               int32   `json:"id"`
	Category         *int32  `json:"category_id"`
	UserCounterparty *string `json:"user_counterparty"`
}
