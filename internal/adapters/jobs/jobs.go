package jobs

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	SynchronizeConnection(ctx context.Context, connectionID int32) error
	HandleSyncError(ctx context.Context, connectionID int32, syncErr error) error
	UpdateConnectorsList(ctx context.Context) ([]model.Connector, []error)
}

type Jobs struct {
	Service Service
}

func NewJobs(service Service) *Jobs {
	return &Jobs{
		Service: service,
	}
}
