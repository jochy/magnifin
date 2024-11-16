package categories

import (
	"magnifin/internal/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateRule(c *gin.Context) {
	var req createCategoryRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.Repository.CreateCategoryRule(c.Request.Context(), model.CategoryRule{
		CategoryID: req.CategoryID,
		Rule:       req.Keywords,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type createCategoryRuleRequest struct {
	CategoryID int32    `json:"category_id"`
	Keywords   []string `json:"keywords"`
}
