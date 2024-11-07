package connectors

import (
	"log/slog"
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/app/model"
	"math"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Connect(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid connector id"})
		return
	} else if id >= math.MaxInt32 || id <= math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid connector id"})
		return
	}

	var req connectRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connectorID := int32(id) //nolint:gosec

	connectInstruction, err := h.service.Connect(
		c.Request.Context(),
		middlewares.GetUser(c.Request.Context()),
		connectorID,
		&model.ConnectParams{
			SID:        uuid.New(),
			SuccessURL: req.SuccessURL,
			ErrorURL:   req.ErrorURL,
		},
	)
	if err != nil {
		slog.Error("unable to connect to connector: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, connectResponse{
		RedirectURL: connectInstruction.RedirectURL,
	})
}

type connectRequest struct {
	SuccessURL string `json:"success_url"`
	ErrorURL   string `json:"error_url"`
}

type connectResponse struct {
	RedirectURL string `json:"redirect_url"`
}
