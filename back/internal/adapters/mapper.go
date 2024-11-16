package adapters

import (
	"fmt"
	"magnifin/internal/app/model"
	"time"
)

type Mapper struct {
	publicURL string
}

type EnrichedTransaction struct {
	ID                int32                      `json:"id"`
	AccountID         int32                      `json:"aid"`
	BankTransactionID *string                    `json:"bid"`
	Amount            float64                    `json:"a"`
	Currency          string                     `json:"c"`
	Direction         model.TransactionDirection `json:"d"`
	Status            model.TransactionStatus    `json:"s"`
	OperationAt       time.Time                  `json:"at"`

	CounterpartyName    *string `json:"name"`
	CounterpartyAccount *string `json:"acc"`
	Reference           *string `json:"ref"`
	Method              *string `json:"m"`

	CounterpartyLogoURL *string `json:"logo"`
	Category            *int32  `json:"ca"`
}

func NewMapper(publicURL string) *Mapper {
	return &Mapper{
		publicURL: publicURL,
	}
}

func (m *Mapper) ToPublicFormat(t *model.Transaction) *EnrichedTransaction {
	if t.Enrichment == nil {
		// Avoid NPE
		t.Enrichment = &model.TransactionEnrichment{}
	}

	counterpartyName := t.Enrichment.CounterpartyName
	if counterpartyName == nil || *counterpartyName == "" {
		counterpartyName = t.CounterpartyName
	}

	if t.Enrichment.UserCounterpartyName != nil && *t.Enrichment.UserCounterpartyName != "" {
		counterpartyName = t.Enrichment.UserCounterpartyName
	}

	reference := t.Enrichment.Reference
	if reference == nil {
		reference = t.Reference
	}

	var logoURL *string
	if t.Enrichment.CounterpartyLogo != nil {
		u := fmt.Sprintf("%s/v1/images/%s", m.publicURL, *t.Enrichment.CounterpartyLogo)
		logoURL = &u
	}

	return &EnrichedTransaction{
		ID:                t.ID,
		AccountID:         t.AccountID,
		BankTransactionID: t.BankTransactionID,
		Amount:            t.Amount,
		Currency:          t.Currency,
		Direction:         t.Direction,
		Status:            t.Status,
		OperationAt:       t.OperationAt,

		CounterpartyName:    counterpartyName,
		CounterpartyAccount: t.CounterpartyAccount,
		Reference:           reference,
		Method:              t.Enrichment.Method,

		CounterpartyLogoURL: logoURL,
		Category:            t.Enrichment.Category,
	}
}
