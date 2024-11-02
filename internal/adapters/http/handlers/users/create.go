package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHandler struct {
	userService Service
}

func NewCreateHandler(userService Service) *CreateHandler {
	return &CreateHandler{
		userService: userService,
	}
}

func (h *CreateHandler) Handle(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Create(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	}

	token, err := h.userService.GenerateJWT(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}

type createRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
