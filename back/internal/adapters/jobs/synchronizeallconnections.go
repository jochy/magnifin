package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/riverqueue/river"
)

type SynchronizeAllConnectionsInput struct{}

func (u SynchronizeAllConnectionsInput) Kind() string {
	return "SynchronizeAllConnections"
}

func (u SynchronizeAllConnectionsInput) InsertOpts() river.InsertOpts {
	return river.InsertOpts{MaxAttempts: 1}
}

func (j *Jobs) NewSynchronizeAllConnectionsPeriodicJob() *river.PeriodicJob {
	return river.NewPeriodicJob(
		river.PeriodicInterval(12*time.Hour),
		func() (river.JobArgs, *river.InsertOpts) {
			return SynchronizeAllConnectionsInput{}, nil
		},
		&river.PeriodicJobOpts{RunOnStart: true},
	)
}

type SynchronizeAllConnectionsWorker struct {
	river.WorkerDefaults[SynchronizeAllConnectionsInput]

	connectionsRepository ConnectionsRepository
	job                   *Jobs
}

func (j *Jobs) NewSynchronizeAllConnectionsWorker() *SynchronizeAllConnectionsWorker {
	return &SynchronizeAllConnectionsWorker{
		connectionsRepository: j.ConnectionsRepository,
		job:                   j,
	}
}

func (w *SynchronizeAllConnectionsWorker) Work(ctx context.Context, _ *river.Job[SynchronizeAllConnectionsInput]) error {
	slog.Info("SynchronizeAllConnectionsWorker started")
	connections, err := w.connectionsRepository.ListConnectionsToSync(ctx)
	if err != nil {
		return err
	}

	for _, connection := range connections {
		err := w.job.Scheduler.Trigger(ctx, SynchronizeConnectionInput{ConnectionID: connection.ID})
		if err != nil {
			return err
		}
	}

	return nil
}
