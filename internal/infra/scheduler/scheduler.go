package scheduler

import (
	"context"
	"magnifin/internal/adapters/jobs"
	"magnifin/internal/infra/database"

	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type Client interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Trigger(ctx context.Context, job river.JobArgs) error
}

type sched struct {
	workers *river.Workers
	client  *river.Client[pgx.Tx]
}

func NewScheduler(db database.Service, jobs *jobs.Jobs) (Client, error) {
	workers := river.NewWorkers()
	addWorkers(workers, jobs)

	riverClient, err := river.NewClient(riverpgxv5.New(db.PgxPool()), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers:      workers,
		PeriodicJobs: periodicJobs(jobs),
	})
	if err != nil {
		return nil, err
	}

	return &sched{
		workers: workers,
		client:  riverClient,
	}, nil
}

func (s *sched) Start(ctx context.Context) error {
	return s.client.Start(ctx)
}

func (s *sched) Stop(ctx context.Context) error {
	return s.client.Stop(ctx)
}

func (s *sched) Trigger(ctx context.Context, job river.JobArgs) error {
	_, err := s.client.Insert(ctx, job, nil)
	if err != nil {
		return err
	}
	return nil
}
