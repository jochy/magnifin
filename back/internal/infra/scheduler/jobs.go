package scheduler

import (
	"magnifin/internal/adapters/jobs"

	"github.com/riverqueue/river"
)

func addWorkers(workers *river.Workers, jobs *jobs.Jobs) {
	river.AddWorker(workers, jobs.NewSynchronizeConnectionWorker())
	river.AddWorker(workers, jobs.NewUpdateConnectorsWorker())
	river.AddWorker(workers, jobs.NewSynchronizeAllConnectionsWorker())
}

func periodicJobs(jobs *jobs.Jobs) []*river.PeriodicJob {
	return []*river.PeriodicJob{
		jobs.NewUpdateConnectorsPeriodicJob(),
		jobs.NewSynchronizeAllConnectionsPeriodicJob(),
	}
}
