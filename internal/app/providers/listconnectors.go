package providers

import (
	"context"
	"magnifin/internal/app/model"

	"golang.org/x/sync/errgroup"
)

func (s *ProviderService) UpdateConnectorsList(ctx context.Context) ([]model.Connector, []error) {
	group, cctx := errgroup.WithContext(ctx)
	group.SetLimit(50)

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

	return updatedConnectors, errors
}
