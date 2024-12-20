package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"magnifin/internal/adapters"
	enricher2 "magnifin/internal/adapters/enricher"
	"magnifin/internal/adapters/http/handlers"
	categorieshandlers "magnifin/internal/adapters/http/handlers/categories"
	connectionshandlers "magnifin/internal/adapters/http/handlers/connections"
	connectorshandlers "magnifin/internal/adapters/http/handlers/connectors"
	imageshandlers "magnifin/internal/adapters/http/handlers/images"
	"magnifin/internal/adapters/http/handlers/providers"
	transactionshandlers "magnifin/internal/adapters/http/handlers/transactions"
	usershandlers "magnifin/internal/adapters/http/handlers/users"
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/adapters/jobs"
	"magnifin/internal/adapters/notifier"
	"magnifin/internal/adapters/providers/gocardless"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/adapters/repository/accounts"
	"magnifin/internal/adapters/repository/categories"
	"magnifin/internal/adapters/repository/connections"
	"magnifin/internal/adapters/repository/connectors"
	"magnifin/internal/adapters/repository/images"
	providersrepo "magnifin/internal/adapters/repository/providers"
	"magnifin/internal/adapters/repository/providerusers"
	"magnifin/internal/adapters/repository/redirect_sessions"
	"magnifin/internal/adapters/repository/transactions"
	usersrepo "magnifin/internal/adapters/repository/users"
	"magnifin/internal/app"
	connections3 "magnifin/internal/app/connections"
	connectors2 "magnifin/internal/app/connectors"
	providers2 "magnifin/internal/app/providers"
	transactions2 "magnifin/internal/app/transactions"
	"magnifin/internal/infra/database"
	scheduler2 "magnifin/internal/infra/scheduler"
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

	openAIKey := os.Getenv("OPENAI_KEY")
	openAIUrl := os.Getenv("OPENAI_BASE_URL")
	openAIModel := os.Getenv("OPENAI_MODEL")
	if (openAIKey == "" && openAIUrl == "") || openAIModel == "" {
		slog.Warn("OPENAI_KEY not set, some features may not work. Please set either OPENAI_KEY, OPENAI_BASE_URL and OPENAI_MODEL")
	}

	publicTransactionMapper := adapters.NewMapper(publicURL)

	// Notifier
	notifier := notifier.NewNotifier(publicTransactionMapper)
	defer notifier.Close() //nolint:errcheck

	// Db
	db := database.NewService()
	userRepository := usersrepo.NewRepository(db, cypherKey)
	providerRepository := providersrepo.NewRepository(db, cypherKey)
	connectorsRepository := connectors.NewRepository(db)
	providerUserRepository := providerusers.NewRepository(db)
	connectionsRepository := connections.NewRepository(db)
	redirectionSessionsRepository := redirect_sessions.NewRepository(db)
	accountsRepository := accounts.NewRepository(db)
	transactionsRepository := transactions.NewRepository(db)
	categoriesRepository := categories.NewRepository(db)
	imagesRepository := images.NewRepository(db)

	// Ports
	providerPorts := []providers2.ProviderPort{
		gocardless.NewGoCardless(publicURL),
	}
	enricher := enricher2.NewEnricher(openAIUrl, openAIKey, openAIModel)

	// Services
	userService := app.NewUserService(userRepository, signKey)
	providerService := providers2.NewProviderService(
		providerRepository,
		connectorsRepository,
		providerUserRepository,
		connectionsRepository,
		redirectionSessionsRepository,
		accountsRepository,
		transactionsRepository,
		userRepository,
		providerPorts,
	)
	connectorsService := connectors2.NewConnectorService(connectorsRepository, providerService)
	connectionsService := connections3.NewConnectionsService(
		connectionsRepository,
		accountsRepository,
		connectorsRepository,
		transactionsRepository,
		providerService,
	)
	transactionsService := transactions2.NewTransactionsService(
		transactionsRepository,
		categoriesRepository,
		imagesRepository,
		userRepository,
		enricher,
		notifier,
	)

	scheduler, err := scheduler2.NewScheduler(db, jobs.NewJobs(providerService, transactionsService, connectionsRepository))
	if err != nil {
		panic(fmt.Sprintf("failed to create scheduler: %s", err))
	}

	// Start the scheduler
	err = scheduler.Start(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to start scheduler: %s", err))
	}

	/*
		trs, err := transactionsRepository.GetAllByUserBetweenDates(context.Background(), &model.User{ID: 8}, time.Now().Add(-36000*time.Hour), time.Now())
		for _, tr := range trs {
			err = scheduler.Trigger(context.Background(), jobs.TransactionEnrichInput{TransactionID: tr.ID})
			if err != nil {
				panic(fmt.Sprintf("failed to trigger transaction enrichment job: %s", err))
			}
		}
	*/

	// Server
	server := server.NewServer(
		handlers.NewHealthHandler(db),
		usershandlers.NewHandler(userService),
		middlewares.NewAuthMiddleware(userService),
		providers.NewHandler(providerService),
		connectorshandlers.NewHandler(connectorsService),
		connectionshandlers.NewHandlers(connectionsService),
		transactionshandlers.NewHandlers(transactionsService, publicTransactionMapper),
		categorieshandlers.NewHandlers(categoriesRepository),
		imageshandlers.NewHandlers(imagesRepository),
		handlers.NewWSHandler(notifier),
	)

	err = providerService.LoadProviderConfigurations()
	if err != nil {
		panic(fmt.Sprintf("failed to load provider configurations: %s", err))
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done

	// Stop the scheduler
	err = scheduler.Stop(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to stop scheduler: %s", err))
	}

	log.Println("Graceful shutdown complete.")
}
