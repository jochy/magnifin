package jobs

import (
	"context"
	"log/slog"

	"github.com/riverqueue/river"
)

type TransactionEnrichInput struct {
	TransactionID int32
}

func (u TransactionEnrichInput) Kind() string {
	return "EnrichTransaction"
}

func (u TransactionEnrichInput) InsertOpts() river.InsertOpts {
	return river.InsertOpts{Priority: 4}
}

type TransactionEnrichWorker struct {
	river.WorkerDefaults[TransactionEnrichInput]

	transactionService TransactionService
}

func (j *Jobs) NewTransactionEnrichWorker() *TransactionEnrichWorker {
	return &TransactionEnrichWorker{
		transactionService: j.TransactionService,
	}
}

func (w *TransactionEnrichWorker) Work(ctx context.Context, input *river.Job[TransactionEnrichInput]) error {
	slog.Debug("TransactionEnrichWorker started")
	return w.transactionService.EnrichTransaction(ctx, input.Args.TransactionID)
}
