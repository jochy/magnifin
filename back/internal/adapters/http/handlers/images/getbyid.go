package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetByID(c *gin.Context) {
	id := c.Param("id")

	image, err := h.ImageRepository.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if image == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.Data(http.StatusOK, image.ContentType, []byte(image.Content))
}
