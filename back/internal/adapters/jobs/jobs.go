package jobs

import (
	"context"
	"magnifin/internal/app/model"

	"github.com/riverqueue/river"
)

type Service interface {
	SynchronizeConnection(ctx context.Context, connectionID int32) error
	HandleSyncError(ctx context.Context, connectionID int32, syncErr error) error
	UpdateConnectorsList(ctx context.Context) ([]model.Connector, []error)
}

type ConnectionsRepository interface {
	ListConnectionsToSync(ctx context.Context) ([]model.Connection, error)
}

type Scheduler interface {
	Trigger(ctx context.Context, job river.JobArgs) error
}

type Jobs struct {
	Service               Service
	ConnectionsRepository ConnectionsRepository
	Scheduler             Scheduler
}

func NewJobs(service Service, connectionsRepository ConnectionsRepository) *Jobs {
	return &Jobs{
		Service:               service,
		ConnectionsRepository: connectionsRepository,
	}
}

func (j *Jobs) SetScheduler(scheduler Scheduler) {
	j.Scheduler = scheduler
}
