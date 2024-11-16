package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() http.Handler {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health", "/v1/ping", "/v1/images/"),
		gin.Recovery(),
	)

	// Public routes
	r.GET("/health", s.healthHandler.HealthHandler)
	r.POST("/v1/login", s.usersHandlers.Login)
	r.POST("/v1/users", s.usersHandlers.Create)
	r.GET("/v1/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/v1/images/:id", s.imagesHandlers.GetByID)

	// Authenticated routes
	auth := r.Group("/v1", s.authMiddleware.Authenticate)

	auth.GET("/check-login", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, nil)
	})
	auth.GET("/ws", s.wsHandler.Listen)

	auth.GET("/providers", s.providersHandlers.List)

	auth.GET("/connectors", s.connectorsHandlers.SearchByName)
	auth.POST("/connectors/:id/connect", s.connectorsHandlers.Connect)

	// User's data
	auth.GET("/connections", s.connectionsHandlers.List)
	auth.DELETE("/connections/:id", s.connectionsHandlers.Delete)

	auth.GET("/transactions", s.transactionsHandler.List)
	auth.GET("/transactions/minmax", s.transactionsHandler.MinMax)
	auth.PATCH("/transactions/:id", s.transactionsHandler.Update)

	auth.GET("/categories", s.categoriesHandlers.List)
	auth.POST("/categories/:id/rule", s.categoriesHandlers.CreateRule)

	// Provider callbacks - No auth because webview on browser / desktop app, so no auth context
	r.GET("/v1/providers/gocardless/callback", s.connectorsHandlers.GoCardlessCallback)

	return r
}
