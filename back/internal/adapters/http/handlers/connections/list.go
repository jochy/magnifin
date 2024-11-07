package connections

import (
	"magnifin/internal/adapters/http/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) List(c *gin.Context) {
	cnxs, err := h.Service.ListConnections(c.Request.Context(), middlewares.GetUser(c.Request.Context()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	connections := make([]connectionWithAccounts, len(cnxs))
	for i, cnx := range cnxs {
		accounts := make([]account, len(cnx.Accounts))
		for j, a := range cnx.Accounts {
			accounts[j] = account{
				ID:            a.ID,
				BankAccountID: a.BankAccountID,
				Name:          a.Name,
				Type:          a.Type,
				Currency:      a.Currency,
				AccountNumber: a.AccountNumber,
				Balance:       a.Balance,
			}
		}
		connections[i] = connectionWithAccounts{
			ID: cnx.Connection.ID,

			Status:             string(cnx.Connection.Status),
			RenewConsentBefore: cnx.Connection.RenewConsentBefore,
			ErrorMessage:       cnx.Connection.ErrorMessage,
			LastSuccessfulSync: cnx.Connection.LastSuccessfulSync,

			Accounts: accounts,
			Connector: connector{
				ID:      cnx.Connector.ID,
				Name:    cnx.Connector.Name,
				LogoURL: cnx.Connector.LogoURL,
			},
		}
	}

	c.JSON(http.StatusOK, connectionsWithAccountsResponse{Connections: connections})
}

type connectionsWithAccountsResponse struct {
	Connections []connectionWithAccounts `json:"connections"`
}

type connectionWithAccounts struct {
	ID int32 `json:"id"`

	Status             string     `json:"status"`
	RenewConsentBefore *time.Time `json:"renew_consent_before"`
	ErrorMessage       *string    `json:"error_message"`
	LastSuccessfulSync *time.Time `json:"last_successful_sync"`

	Accounts  []account `json:"accounts"`
	Connector connector `json:"connector"`
}

type account struct {
	ID            int32   `json:"id"`
	BankAccountID *string `json:"bank_account_id"`
	Name          *string `json:"name"`
	Type          *string `json:"type"`
	Currency      *string `json:"currency"`
	AccountNumber *string `json:"account_number"`
	Balance       float64 `json:"balance"`
}

type connector struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}
