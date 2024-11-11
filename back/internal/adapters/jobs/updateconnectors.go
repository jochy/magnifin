package jobs

import (
	"context"
	"time"

	"github.com/riverqueue/river"
)

type UpdateConnectorsInput struct{}

func (u UpdateConnectorsInput) Kind() string {
	return "UpdateConnectors"
}

func (u UpdateConnectorsInput) InsertOpts() river.InsertOpts {
	return river.InsertOpts{MaxAttempts: 5, Priority: 4}
}

func (j *Jobs) NewUpdateConnectorsPeriodicJob() *river.PeriodicJob {
	return river.NewPeriodicJob(
		river.PeriodicInterval(30*time.Minute),
		func() (river.JobArgs, *river.InsertOpts) {
			return UpdateConnectorsInput{}, nil
		},
		&river.PeriodicJobOpts{RunOnStart: true},
	)
}

type UpdateConnectorsWorker struct {
	river.WorkerDefaults[UpdateConnectorsInput]

	service Service
}

func (j *Jobs) NewUpdateConnectorsWorker() *UpdateConnectorsWorker {
	return &UpdateConnectorsWorker{
		service: j.Service,
	}
}

func (w *UpdateConnectorsWorker) Work(ctx context.Context, _ *river.Job[UpdateConnectorsInput]) error {
	_, err := w.service.UpdateConnectorsList(ctx)
	if len(err) > 0 {
		return err[0]
	}

	return nil
}
