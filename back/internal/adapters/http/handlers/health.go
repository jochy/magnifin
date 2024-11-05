package handlers

import (
	"github.com/gin-gonic/gin"
)

type Db interface {
	Health() map[string]string
}

type HealthHandler struct {
	Db Db
}

func NewHealthHandler(db Db) *HealthHandler {
	return &HealthHandler{
		Db: db,
	}
}

func (h *HealthHandler) HealthHandler(c *gin.Context) {
	details := h.Db.Health()
	if details == nil {
		c.JSON(500, nil)
		return
	}

	if details["status"] != "up" {
		c.JSON(500, details)
		return
	}

	c.JSON(200, details)
}
