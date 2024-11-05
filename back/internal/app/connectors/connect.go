package connectors

import (
	"context"
	"errors"
	"magnifin/internal/app/model"
)

func (s *Service) Connect(ctx context.Context, user *model.User, connectorID int32, params *model.ConnectParams) (*model.ConnectInstruction, error) {
	connector, err := s.repository.GetByID(ctx, connectorID)
	if err != nil {
		return nil, err
	} else if connector == nil {
		return nil, errors.New("connector not found in db")
	}

	return s.providerService.Connect(ctx, user, connector, params)
}
