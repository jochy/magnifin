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

	// Authenticated routes
	auth := r.Group("/v1", s.authMiddleware.Authenticate)

	auth.GET("/check-login", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, nil)
	})

	auth.GET("/providers", s.providersHandlers.List)
	auth.POST("/providers/:id", s.providersHandlers.Update)

	auth.GET("/connectors", s.connectorsHandlers.SearchByName)
	auth.POST("/connectors/:id/connect", s.connectorsHandlers.Connect)

	// Provider callbacks
	auth.GET("/providers/gocardless/callback", s.connectorsHandlers.GoCardlessCallback)

	return r
}