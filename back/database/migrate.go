package main

import (
	"context"
	"errors"
	"magnifin/internal/infra/database"

	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

func main() {
	driver, err := postgres.WithInstance(database.NewService().Driver(), &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		// Ignore
	} else if err != nil {
		panic(err)
	}

	riverdriver := riverpgxv5.New(database.NewService().PgxPool())
	migrator, err := rivermigrate.New(riverdriver, nil)
	if err != nil {
		panic(err)
	}

	_, err = migrator.Migrate(context.Background(), rivermigrate.DirectionUp, nil)
	if err != nil {
		panic(err)
	}
}
