package middlewares

import (
	"context"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"

	"github.com/gin-gonic/gin"
)

const UserContextKey = "user-auth-middleware"

type UserService interface {
	FromJWT(ctx context.Context, token string) (*model.User, error)
}

type AuthMiddleware struct {
	userService UserService
}

func NewAuthMiddleware(userService UserService) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
	}
}

func (m *AuthMiddleware) Authenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	tokenQuery := ctx.Query("token")
	if token == "" && tokenQuery == "" {
		slog.Debug("missing Authorization header")
		ctx.JSON(401, gin.H{"error": "missing Authorization header"})
		ctx.Abort()
		return
	}

	tokenToUse := token
	if tokenToUse == "" {
		tokenToUse = tokenQuery
	}

	user, err := m.userService.FromJWT(ctx.Request.Context(), tokenToUse)
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing token: %v", err))
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if user == nil {
		slog.Debug("invalid token")
		ctx.JSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	slog.Debug(fmt.Sprintf("authenticated user: %v", user))
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), UserContextKey, user))

	ctx.Next()
}

func GetUser(ctx context.Context) *model.User {
	user, ok := ctx.Value(UserContextKey).(*model.User)
	if !ok {
		return nil
	}

	return user
}
