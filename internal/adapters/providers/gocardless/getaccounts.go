package gocardless

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"
	"net/http"
	"strconv"
)

const (
	goCardlessAccountsUpdate = "/api/v2/accounts/%s/"
	goCardlessAccounts       = "/api/v2/accounts/%s/details/"
	goCardlessBalance        = "/api/v2/accounts/%s/balances/"
)

func (g *GoCardless) GetAccounts(
	ctx context.Context,
	provider *model.Provider,
	_ *model.ProviderUser,
	connection *model.Connection,
) ([]model.Account, error) {
	accountIDs, err := g.getAccountIDs(ctx, provider, connection)
	if err != nil {
		return nil, err
	}

	// For each account, retrieve the data
	accounts := make([]model.Account, len(accountIDs))
	for i, accountID := range accountIDs {
		account, err := g.getAccountByID(ctx, provider, connection, accountID)
		if err != nil {
			return nil, err
		}

		accounts[i] = *account
	}

	return accounts, nil
}

func (g *GoCardless) triggerAccountSync(ctx context.Context, provider *model.Provider, _ *model.Connection, accountID string) error {
	req, err := g.newRequest(ctx, provider, http.MethodGet, fmt.Sprintf(goCardlessAccountsUpdate, accountID), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to get sync account: " + resp.Status)
	}

	slog.Info("Account sync triggered for account: " + accountID)
	return nil
}

func (g *GoCardless) getAccountIDs(ctx context.Context, provider *model.Provider, connection *model.Connection) ([]string, error) {
	req, err := g.newRequest(ctx, provider, http.MethodGet, goCardlessRequisition+connection.ProviderConnectionID+"/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get connection: " + resp.Status)
	}

	var res requisitionResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res.Accounts, nil
}

func (g *GoCardless) getAccountByID(ctx context.Context, provider *model.Provider, connection *model.Connection, accountID string) (*model.Account, error) {
	err := g.triggerAccountSync(ctx, provider, connection, accountID)
	if err != nil {
		return nil, err
	}

	req, err := g.newRequest(ctx, provider, http.MethodGet, fmt.Sprintf(goCardlessAccounts, accountID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get account: " + resp.Status)
	}

	var res accountMetadata
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	req, err = g.newRequest(ctx, provider, http.MethodGet, fmt.Sprintf(goCardlessBalance, accountID), nil)
	if err != nil {
		return nil, err
	}

	respBalance, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer respBalance.Body.Close() //nolint: errcheck

	if respBalance.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get account balance: " + respBalance.Status)
	}

	var resBalance accountBalance
	if err := json.NewDecoder(respBalance.Body).Decode(&resBalance); err != nil {
		return nil, err
	}

	return &model.Account{
		ConnectionID:      connection.ID,
		ProviderAccountID: accountID,
		Name:              res.Account.Name,
		Type:              res.Account.Product,
		Currency:          res.Account.Currency,
		AccountNumber:     res.Account.Iban,
		Balance:           resBalance.getBalance(),
	}, nil
}

type accountMetadata struct {
	Account accountData `json:"account"`
}

type accountData struct {
	Iban     *string `json:"iban"`
	Currency *string `json:"currency"`
	Name     *string `json:"name"`
	Product  *string `json:"product"`
}

type accountBalance struct {
	Balances []balance `json:"balances"`
}

type balance struct {
	BalanceAmount balanceAmount `json:"balanceAmount"`
	BalanceType   string        `json:"balanceType"`
	ReferenceDate string        `json:"referenceDate"`
}

type balanceAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func (b *accountBalance) getBalance() float64 {
	for _, balance := range b.Balances {
		if balance.BalanceType == "expected" {
			amount, err := strconv.ParseFloat(balance.BalanceAmount.Amount, 64)
			if err != nil {
				slog.Warn("Failed to parse balance amount: " + balance.BalanceAmount.Amount)
				return 0
			}

			return amount
		}
	}

	slog.Warn("No expected balance found")
	return 0
}
