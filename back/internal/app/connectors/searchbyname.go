package connectors

import (
	"context"
	"log/slog"
	"magnifin/internal/app/model"
)

func (s *Service) SearchByName(ctx context.Context, name string) ([]model.Connector, error) {
	res, err := s.repository.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		slog.Debug("No connectors found with fuzzy search, trying like search")
		return s.repository.LikeSearchByName(ctx, name)
	}
	return res, nil
}
