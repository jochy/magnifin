package transactions

import (
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) List(c *gin.Context) {
	user := middlewares.GetUser(c.Request.Context())
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	fromStr := c.Query("from")
	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
		return
	}

	toStr := c.Query("to")
	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
		return
	}

	transactions, err := h.Service.GetAllByUserBetweenDates(c.Request.Context(), user, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t := make([]enrichedTransaction, len(transactions))
	for i, tx := range transactions {
		t[i] = *toPublicFormat(&tx)
	}

	c.JSON(http.StatusOK, listResponse{Transactions: t})
}

type listResponse struct {
	Transactions []enrichedTransaction `json:"transactions"`
}

type enrichedTransaction struct {
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

	CounterpartyLogoURL *string `json:"logo"`
	Category            *string `json:"ca"`
}

func toPublicFormat(t *model.Transaction) *enrichedTransaction {
	if t.Enrichment == nil {
		// Avoid NPE
		t.Enrichment = &model.TransactionEnrichment{}
	}

	counterpartyName := t.Enrichment.CounterpartyName
	if counterpartyName == nil {
		counterpartyName = t.CounterpartyName
	}

	reference := t.Enrichment.Reference
	if reference == nil {
		reference = t.Reference
	}

	return &enrichedTransaction{
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

		CounterpartyLogoURL: t.Enrichment.CounterpartyLogoURL,
		Category:            t.Enrichment.Category,
	}
}