package categories

import (
	"context"
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/adapters/jobs"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/scheduler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateRule(c *gin.Context) {
	var req createCategoryRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rule, err := h.Repository.CreateCategoryRule(c.Request.Context(), model.CategoryRule{
		CategoryID: req.CategoryID,
		Rule:       req.Keywords,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.ApplyToAll {
		_ = scheduler.Scheduler.Trigger(context.Background(), jobs.ApplyCategoryRuleInput{
			RuleID: rule.ID,
			UserID: middlewares.GetUser(c.Request.Context()).ID,
		})
	}

	c.JSON(http.StatusNoContent, nil)
}

type createCategoryRuleRequest struct {
	CategoryID int32    `json:"category_id"`
	Keywords   []string `json:"keywords"`
	ApplyToAll bool     `json:"apply_to_all"`
}
