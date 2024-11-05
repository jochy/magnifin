package providers

import (
	"context"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"

	"golang.org/x/sync/errgroup"
)

const (
	maxConcurrency = 50
)

func (s *ProviderService) UpdateConnectorsList(ctx context.Context) ([]model.Connector, []error) {
	slog.Info("Updating connectors list")

	group, cctx := errgroup.WithContext(ctx)
	group.SetLimit(maxConcurrency)

	var connectors []model.Connector
	var errors []error

	// Retrieve connectors from all ports
	for _, p := range s.ports {
		port := p
		group.Go(func() error {
			provider, err := s.providerRepository.GetByName(cctx, port.Name())

			if err != nil {
				errors = append(errors, err)
				return nil
			} else if provider == nil {
				errors = append(errors, fmt.Errorf("provider %s not found", port.Name()))
				return nil
			}

			if !provider.Enabled {
				return nil
			}

			c, err := port.ListConnectors(cctx, provider)
			if err != nil {
				errors = append(errors, err)
			} else {
				connectors = append(connectors, c...)
			}

			return nil
		})
	}

	err := group.Wait()
	if err != nil {
		errors = append(errors, err)
	}

	// Upsert connectors
	group, cctx = errgroup.WithContext(ctx)
	group.SetLimit(maxConcurrency)

	var updatedConnectors []model.Connector
	for _, c := range connectors {
		group.Go(func() error {
			updated, err := s.connectorRepository.Upsert(cctx, &c)
			if err != nil {
				errors = append(errors, err)
			} else {
				updatedConnectors = append(updatedConnectors, *updated)
			}
			return nil
		})
	}

	err = group.Wait()
	if err != nil {
		errors = append(errors, err)
	}

	slog.Info("Connectors list updated")
	slog.Debug(fmt.Sprintf("Total connectors: %d", len(updatedConnectors)))
	slog.Debug(fmt.Sprintf("Total errors: %s", errors))
	return updatedConnectors, errors
}
