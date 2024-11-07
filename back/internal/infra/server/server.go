package server

import (
	"fmt"
	"magnifin/internal/adapters/http/handlers"
	"magnifin/internal/adapters/http/handlers/connections"
	"magnifin/internal/adapters/http/handlers/connectors"
	"magnifin/internal/adapters/http/handlers/providers"
	"magnifin/internal/adapters/http/handlers/users"
	"magnifin/internal/adapters/http/middlewares"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int

	healthHandler       *handlers.HealthHandler
	usersHandlers       *users.Handler
	authMiddleware      *middlewares.AuthMiddleware
	providersHandlers   *providers.Handler
	connectorsHandlers  *connectors.Handler
	connectionsHandlers *connections.Handlers
}

func NewServer(
	healthHandler *handlers.HealthHandler,
	usersHandlers *users.Handler,
	authMiddleware *middlewares.AuthMiddleware,
	providersHandlers *providers.Handler,
	connectorsHandlers *connectors.Handler,
	connectionsHandlers *connections.Handlers,
) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:                port,
		healthHandler:       healthHandler,
		usersHandlers:       usersHandlers,
		authMiddleware:      authMiddleware,
		providersHandlers:   providersHandlers,
		connectorsHandlers:  connectorsHandlers,
		connectionsHandlers: connectionsHandlers,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.registerRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
