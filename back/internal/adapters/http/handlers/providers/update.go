package providers

import (
	"magnifin/internal/app/model"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(ctx *gin.Context) {
	var req updateProviderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid provider id"})
		return
	} else if id >= math.MaxInt32 || id <= math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid provider id"})
		return
	}

	provider := model.Provider{
		ID:        int32(id), //nolint:gosec
		Name:      req.Name,
		Enabled:   req.Enabled,
		AccessKey: req.AccessKey,
		Secret:    req.Secret,
	}

	updatedProvider, err := h.service.UpdateProvider(ctx, provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, toProviderResponse(*updatedProvider))
}

type updateProviderRequest struct {
	Name      string  `json:"name" required:"true"`
	Enabled   bool    `json:"enabled" required:"true"`
	AccessKey *string `json:"access_key"`
	Secret    *string `json:"secret"`
}
