package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"magnifin/internal/adapters/http/handlers"
	connectorshandlers "magnifin/internal/adapters/http/handlers/connectors"
	"magnifin/internal/adapters/http/handlers/providers"
	usershandlers "magnifin/internal/adapters/http/handlers/users"
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/adapters/providers/gocardless"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/adapters/repository/connectors"
	providersrepo "magnifin/internal/adapters/repository/providers"
	"magnifin/internal/adapters/repository/providerusers"
	usersrepo "magnifin/internal/adapters/repository/users"
	"magnifin/internal/app"
	connectors2 "magnifin/internal/app/connectors"
	providers2 "magnifin/internal/app/providers"
	"magnifin/internal/infra/database"
	"magnifin/internal/infra/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	signKey := os.Getenv("JWT_SIGN_KEY")
	if signKey == "" {
		slog.Warn("JWT_SIGN_KEY not set, using random value. This means that each time the server restarts, all JWT tokens will be invalidated.")
		signKey = uuid.New().String()
	}

	cypherKey := os.Getenv("CYPHER_KEY")
	if cypherKey == "" {
		panic("DB_CYPHER_KEY is required")
	}
	cypherKey = repository.Generate32ByteKey(cypherKey)

	publicURL := os.Getenv("PUBLIC_URL")
	if publicURL == "" {
		panic("PUBLIC_URL is required")
	}

	// Db
	db := database.NewService()
	userRepository := usersrepo.NewRepository(db, cypherKey)
	providerRepository := providersrepo.NewRepository(db, cypherKey)
	connectorsRepository := connectors.NewRepository(db)
	providerUserRepository := providerusers.NewRepository(db)

	// Ports
	providerPorts := []providers2.ProviderPort{
		gocardless.NewGoCardless(publicURL),
	}

	// Services
	userService := app.NewUserService(userRepository, signKey)
	providerService := providers2.NewProviderService(providerRepository, connectorsRepository, providerUserRepository, providerPorts)
	connectorsService := connectors2.NewConnectorService(connectorsRepository, providerService)

	// Refresh the connectors list in backgrounds
	go func() {
		providerService.UpdateConnectorsList(context.Background())
	}()

	// Server
	server := server.NewServer(
		handlers.NewHealthHandler(db),
		usershandlers.NewHandler(userService),
		middlewares.NewAuthMiddleware(userService),
		providers.NewHandler(providerService),
		connectorshandlers.NewHandler(connectorsService),
	)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
