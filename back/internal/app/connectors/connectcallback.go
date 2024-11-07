package connectors

import (
	"context"
	"errors"
)

func (s *Service) ConnectCallback(
	ctx context.Context,
	connectorID int32,
	sid string,
	providerConnectionID *string,
) error {
	connector, err := s.repository.GetByID(ctx, connectorID)
	if err != nil {
		return err
	} else if connector == nil {
		return errors.New("connector not found in db")
	}

	return s.providerService.ConnectCallback(ctx, connector, sid, providerConnectionID)
}
