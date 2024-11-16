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

type TransactionService interface {
	EnrichTransaction(ctx context.Context, transactionID int32) error
	ApplyCategoryRule(ctx context.Context, ruleID int32, userID int32) error
}

type ConnectionsRepository interface {
	ListConnectionsToSync(ctx context.Context) ([]model.Connection, error)
}

type Scheduler interface {
	Trigger(ctx context.Context, job river.JobArgs) error
}

type Jobs struct {
	Service               Service
	TransactionService    TransactionService
	ConnectionsRepository ConnectionsRepository
	Scheduler             Scheduler
}

func NewJobs(service Service, transactionService TransactionService, connectionsRepository ConnectionsRepository) *Jobs {
	return &Jobs{
		Service:               service,
		TransactionService:    transactionService,
		ConnectionsRepository: connectionsRepository,
	}
}

func (j *Jobs) SetScheduler(scheduler Scheduler) {
	j.Scheduler = scheduler
}
