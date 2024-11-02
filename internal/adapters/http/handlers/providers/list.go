package providers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(ctx *gin.Context) {
	providers, err := h.service.List(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("error listing providers: %s", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := listProviderResponse{}
	for _, provider := range providers {
		response.Providers = append(response.Providers, toProviderResponse(provider))
	}

	ctx.JSON(http.StatusOK, response)
}

type listProviderResponse struct {
	Providers []providerResponse `json:"providers"`
}
