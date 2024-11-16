package jobs

import (
	"context"
	"log/slog"

	"github.com/riverqueue/river"
)

type ApplyCategoryRuleInput struct {
	RuleID int32
	UserID int32
}

func (u ApplyCategoryRuleInput) Kind() string {
	return "ApplyCategoryRule"
}

func (u ApplyCategoryRuleInput) InsertOpts() river.InsertOpts {
	return river.InsertOpts{Priority: 3}
}

type ApplyCategoryRuleWorker struct {
	river.WorkerDefaults[ApplyCategoryRuleInput]

	transactionService TransactionService
}

func (j *Jobs) NewApplyCategoryRuleWorker() *ApplyCategoryRuleWorker {
	return &ApplyCategoryRuleWorker{
		transactionService: j.TransactionService,
	}
}

func (w *ApplyCategoryRuleWorker) Work(ctx context.Context, input *river.Job[ApplyCategoryRuleInput]) error {
	slog.Debug("ApplyCategoryRuleWorker started")
	return w.transactionService.ApplyCategoryRule(ctx, input.Args.RuleID, input.Args.UserID)
}
