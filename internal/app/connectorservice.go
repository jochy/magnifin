package app

import (
	"context"
	"log/slog"
	"magnifin/internal/app/model"
)

type Repository interface {
	SearchByName(ctx context.Context, name string) ([]model.Connector, error)
	LikeSearchByName(ctx context.Context, name string) ([]model.Connector, error)
	GetByID(ctx context.Context, id int32) (*model.Connector, error)
}

type Service struct {
	repository Repository
}

func NewConnectorService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

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

func (s *Service) GetByID(ctx context.Context, id int32) (*model.Connector, error) {
	return s.repository.GetByID(ctx, id)
}
