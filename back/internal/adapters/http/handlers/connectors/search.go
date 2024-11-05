package connectors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchByName(c *gin.Context) {
	name := c.Query("name")

	connectors, err := h.service.SearchByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	connectorsResponse := make([]connectorResponse, len(connectors))
	for i, c := range connectors {
		connectorsResponse[i] = connectorResponse{
			ID:         c.ID,
			Name:       c.Name,
			LogoURL:    c.LogoURL,
			ProviderID: c.ProviderID,
		}
	}

	c.JSON(http.StatusOK, searchResponse{Connectors: connectorsResponse})
}

type searchResponse struct {
	Connectors []connectorResponse `json:"connectors"`
}

type connectorResponse struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LogoURL    string `json:"logo_url"`
	ProviderID int32  `json:"provider_id"`
}
