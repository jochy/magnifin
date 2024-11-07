package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() http.Handler {
	r := gin.Default()

	// Public routes
	r.GET("/health", s.healthHandler.HealthHandler)
	r.POST("/v1/login", s.usersHandlers.Login)
	r.POST("/v1/users", s.usersHandlers.Create)
	r.GET("/v1/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Authenticated routes
	auth := r.Group("/v1", s.authMiddleware.Authenticate)

	auth.GET("/check-login", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, nil)
	})

	auth.GET("/providers", s.providersHandlers.List)
	auth.POST("/providers/:id", s.providersHandlers.Update)

	auth.GET("/connectors", s.connectorsHandlers.SearchByName)
	auth.POST("/connectors/:id/connect", s.connectorsHandlers.Connect)

	// User's data
	auth.GET("/connections", s.connectionsHandlers.List)
	auth.DELETE("/connections/:id", s.connectionsHandlers.Delete)

	// Provider callbacks - No auth because webview on browser / desktop app, so no auth context
	r.GET("/v1/providers/gocardless/callback", s.connectorsHandlers.GoCardlessCallback)

	return r
}
