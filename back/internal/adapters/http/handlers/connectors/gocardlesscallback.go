package connectors

import (
	"fmt"
	"log/slog"
	"magnifin/internal/adapters/http/middlewares"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GoCardlessCallback(c *gin.Context) {
	connectionID := c.Query("ref")
	connectorID := c.Query("c")
	sid := c.Query("sid")
	successURL := c.Query("s")
	errorURL := c.Query("e")

	errorCode := c.Query("error")
	if errorCode != "" {
		slog.Warn(fmt.Sprintf("error code: %s", errorCode))
		c.Redirect(http.StatusTemporaryRedirect, errorURL+"?msg="+c.Query("details"))
		return
	}

	if connectorID == "" {
		slog.Warn("connectorID is empty")
		c.Redirect(http.StatusTemporaryRedirect, errorURL)
		return
	}
	connectorIDInt, err := strconv.Atoi(connectorID)
	if err != nil {
		slog.Warn(fmt.Sprintf("failed to convert connectorID to int: %s", err))
		c.Redirect(http.StatusTemporaryRedirect, errorURL)
		return
	}

	if connectorIDInt > math.MaxInt32 || connectorIDInt < math.MinInt32 {
		slog.Warn("connectorID is out of range")
		c.Redirect(http.StatusTemporaryRedirect, errorURL)
		return
	}

	cID := int32(connectorIDInt) //nolint:gosec

	err = h.service.ConnectCallback(c.Request.Context(), middlewares.GetUser(c.Request.Context()), cID, sid, &connectionID)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to connect callback: %s", err))

		c.Redirect(http.StatusTemporaryRedirect, errorURL)
		return
	}

	slog.Info("connection successfully created")
	c.Redirect(http.StatusTemporaryRedirect, successURL)
}
