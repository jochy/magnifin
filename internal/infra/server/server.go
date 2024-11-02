package server

import (
	"fmt"
	"magnifin/internal/adapters/http/handlers"
	"magnifin/internal/adapters/http/handlers/users"
	"magnifin/internal/adapters/http/middlewares"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int

	healthHandler     *handlers.HealthHandler
	loginHandler      *users.LoginHandler
	createUserHandler *users.CreateHandler

	authMiddleware *middlewares.AuthMiddleware
}

func NewServer(
	healthHandler *handlers.HealthHandler,
	loginHandler *users.LoginHandler,
	createHandler *users.CreateHandler,
	authMiddleware *middlewares.AuthMiddleware,
) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:              port,
		healthHandler:     healthHandler,
		loginHandler:      loginHandler,
		createUserHandler: createHandler,
		authMiddleware:    authMiddleware,
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
