package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	userService Service
}

func NewLoginHandler(userService Service) *LoginHandler {
	return &LoginHandler{
		userService: userService,
	}
}

func (h *LoginHandler) Handle(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	} else if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := h.userService.GenerateJWT(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"internal_error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: token})
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
