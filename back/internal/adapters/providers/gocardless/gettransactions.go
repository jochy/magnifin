package gocardless

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	goCardlessTransactions = "/api/v2/accounts/%s/transactions/"
)

func (g *GoCardless) GetTransactions(
	ctx context.Context,
	provider *model.Provider,
	_ *model.ProviderUser,
	_ *model.Connection,
	account *model.Account,
) ([]model.Transaction, error) {
	req, err := g.newRequest(ctx, provider, http.MethodGet, fmt.Sprintf(goCardlessTransactions, account.ProviderAccountID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, model.ErrRateLimited
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get transactions: %s", resp.Status)
	}

	var transactionsResp transactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionsResp)
	if err != nil {
		return nil, err
	}

	// yes, I should really prealloc, but... I'm lazy
	var transactions []model.Transaction //nolint: prealloc
	for _, transaction := range transactionsResp.Transactions.Booked {
		transactions = append(transactions, transaction.toDomain(account.ID, true))
	}
	for _, transaction := range transactionsResp.Transactions.Pending {
		transactions = append(transactions, transaction.toDomain(account.ID, false))
	}

	return transactions, nil
}

type transactionResponse struct {
	Transactions transactionObjectResponse `json:"transactions"`
}

type transactionObjectResponse struct {
	Booked  []goCardlessTransaction `json:"booked"`
	Pending []goCardlessTransaction `json:"pending"`
}

type goCardlessTransaction struct {
	TransactionId         *string `json:"transactionId"`
	InternalTransactionId *string `json:"internalTransactionId"`

	BookingDate     *string                        `json:"bookingDate"`
	BookingDateTime *time.Time                     `json:"bookingDateTime"`
	ValueDate       *string                        `json:"valueDate"`
	ValueDateTime   *time.Time                     `json:"valueDateTime"`
	Amount          goCardlessTransactionAmount    `json:"transactionAmount"`
	DebtorName      *string                        `json:"debtorName"`
	CreditorName    *string                        `json:"creditorName"`
	DebtorAccount   *goCardlessCounterPartyAccount `json:"debtorAccount"`
	CreditorAccount *goCardlessCounterPartyAccount `json:"creditorAccount"`
	Reference       *string                        `json:"remittanceInformationUnstructured"`
	References      []string                       `json:"remittanceInformationUnstructuredArray"`
}

type goCardlessTransactionAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type goCardlessCounterPartyAccount struct {
	Iban      *string `json:"iban"`
	Bban      *string `json:"bban"`
	Pan       *string `json:"pan"`
	MaskedPan *string `json:"maskedPan"`
	Msisdn    *string `json:"msisdn"`
}

func (transaction *goCardlessTransaction) toDomain(accountID int32, isBooked bool) model.Transaction {
	providerTransID := transaction.TransactionId
	if providerTransID == nil {
		providerTransID = transaction.InternalTransactionId
		if providerTransID == nil {
			slog.Warn("transaction without ID, generating one")
			id := uuid.New().String()
			providerTransID = &id
		}
	}

	transactionStatus := model.TransactionStatusPending
	if isBooked {
		transactionStatus = model.TransactionStatusCompleted
	}

	var counterpartyName *string
	var counterpartyAccount *string
	if transaction.DebtorName != nil {
		counterpartyName = transaction.DebtorName

		if transaction.DebtorAccount != nil {
			counterpartyAccount = transaction.DebtorAccount.getIdentifier()
		}
	} else if transaction.CreditorName != nil {
		counterpartyName = transaction.CreditorName

		if transaction.CreditorAccount != nil {
			counterpartyAccount = transaction.CreditorAccount.getIdentifier()
		}
	}

	reference := transaction.Reference
	if reference == nil && len(transaction.References) > 0 {
		ref := strings.Join(transaction.References, ", ")
		reference = &ref
	}

	if reference != nil {
		reference = cleanString(reference)
	}

	if counterpartyName == nil {
		counterpartyName = reference
	} else {
		counterpartyName = cleanString(counterpartyName)
	}

	var operationDate time.Time
	if transaction.ValueDateTime != nil { //nolint:gocritic
		operationDate = *transaction.ValueDateTime
	} else if transaction.ValueDate != nil {
		operationDate, _ = time.Parse("2006-01-02", *transaction.ValueDate)
	} else if transaction.BookingDateTime != nil {
		operationDate = *transaction.BookingDateTime
	} else if transaction.BookingDate != nil {
		operationDate, _ = time.Parse("2006-01-02", *transaction.BookingDate)
	} else {
		slog.Warn("transaction without date, using current time, even though it's wrong")
		operationDate = time.Now()
	}

	return model.Transaction{
		AccountID:             accountID,
		ProviderTransactionID: *providerTransID,
		BankTransactionID:     transaction.TransactionId,
		Amount:                transaction.Amount.getAmount(),
		Currency:              transaction.Amount.Currency,
		Direction:             transaction.Amount.getDirection(),
		Status:                transactionStatus,
		OperationAt:           operationDate,
		CounterpartyName:      counterpartyName,
		CounterpartyAccount:   counterpartyAccount,
		Reference:             reference,
	}
}

func cleanString(reference *string) *string {
	if reference == nil {
		return nil
	}

	charToRemove := []string{"\n", "\r", "\t", "#", "@", "(", ")", "-", "_", ",", "  "}
	for _, char := range charToRemove {
		*reference = strings.ReplaceAll(*reference, char, " ")
	}

	res := strings.TrimSpace(*reference)
	return &res
}

func (a *goCardlessTransactionAmount) getAmount() float64 {
	amount, err := strconv.ParseFloat(a.Amount, 64)
	if err != nil {
		slog.Warn("Failed to parse amount: " + a.Amount)
		return 0
	}
	return math.Abs(amount)
}

func (a *goCardlessTransactionAmount) getDirection() model.TransactionDirection {
	amount, err := strconv.ParseFloat(a.Amount, 64)
	if err != nil {
		slog.Warn("Failed to parse amount: " + a.Amount)
		amount = 0
	}
	if amount > 0 {
		return model.TransactionDirectionCredit
	}
	return model.TransactionDirectionDebit
}

func (a *goCardlessCounterPartyAccount) getIdentifier() *string {
	list := []*string{a.Iban, a.Bban, a.Pan, a.MaskedPan, a.Msisdn}
	for _, id := range list {
		if id != nil {
			return id
		}
	}

	return nil
}
