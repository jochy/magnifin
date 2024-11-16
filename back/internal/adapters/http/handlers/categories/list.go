package categories

import (
	"magnifin/internal/adapters/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) List(c *gin.Context) {
	user := middlewares.GetUser(c.Request.Context())
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	categories, err := h.Repository.GetAllCategoriesByUserID(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cats := make([]categoryResponse, len(categories))
	for i, cat := range categories {
		cats[i] = categoryResponse{
			ID:              cat.ID,
			Name:            cat.Name,
			UserID:          cat.UserID,
			Color:           cat.Color,
			Icon:            cat.Icon,
			IncludeInBudget: cat.IncludeInBudget,
		}
	}

	c.JSON(http.StatusOK, listResponse{Categories: cats})
}

type listResponse struct {
	Categories []categoryResponse `json:"categories"`
}

type categoryResponse struct {
	ID              int32  `json:"id"`
	Name            string `json:"name"`
	UserID          *int32 `json:"uid"`
	Color           string `json:"color"`
	Icon            string `json:"icon"`
	IncludeInBudget bool   `json:"include_in_budget"`
}
