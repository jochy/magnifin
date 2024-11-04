package connectors

import "github.com/gin-gonic/gin"

func (h *Handler) GoCardlessCallback(c *gin.Context) {
	ref := c.Query("ref")
	c.JSON(200, gin.H{"message": ref})
}
