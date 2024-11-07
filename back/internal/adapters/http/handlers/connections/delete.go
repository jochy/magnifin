package connections

import (
	"magnifin/internal/adapters/http/middlewares"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	// Convert the id to int32
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if idInt > math.MaxInt32 || idInt < math.MinInt32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.Service.DeleteConnection(c, middlewares.GetUser(c.Request.Context()), int32(idInt)) //nolint:gosec
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
