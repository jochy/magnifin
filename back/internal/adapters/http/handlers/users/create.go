package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	}

	token, err := h.service.GenerateJWT(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}

type createRequest struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}
