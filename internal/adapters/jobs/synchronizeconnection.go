package jobs

import (
	"context"

	"github.com/riverqueue/river"
)

type SynchronizeConnectionInput struct {
	ConnectionID int32
}

func (s SynchronizeConnectionInput) Kind() string {
	return "SynchronizeConnection"
}

func (s SynchronizeConnectionInput) InsertOpts() river.InsertOpts {
	return river.InsertOpts{MaxAttempts: 2}
}

type SynchronizeConnectionWorker struct {
	river.WorkerDefaults[SynchronizeConnectionInput]

	service Service
}

func (j *Jobs) NewSynchronizeConnectionWorker() *SynchronizeConnectionWorker {
	return &SynchronizeConnectionWorker{
		service: j.Service,
	}
}

func (w *SynchronizeConnectionWorker) Work(ctx context.Context, job *river.Job[SynchronizeConnectionInput]) error {
	err := w.service.SynchronizeConnection(ctx, job.Args.ConnectionID)
	if err != nil {
		err = w.service.HandleSyncError(ctx, job.Args.ConnectionID, err)
	}

	return err
}
