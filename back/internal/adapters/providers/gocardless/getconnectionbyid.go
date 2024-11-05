package gocardless

import (
	"context"
	"encoding/json"
	"errors"
	"magnifin/internal/app/model"
	"net/http"
	"time"
)

const (
	goCardlessAgreement = "/api/v2/agreements/enduser/"
)

func (g *GoCardless) GetConnectionByID(
	ctx context.Context,
	provider *model.Provider,
	providerUser *model.ProviderUser,
	connector *model.Connector,
	connectionID string,
) (*model.Connection, error) {
	req, err := g.newRequest(ctx, provider, http.MethodGet, goCardlessRequisition+connectionID+"/", nil)
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
		return nil, errors.New("failed to get connection: " + resp.Status)
	}

	var res requisitionResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	req, err = g.newRequest(ctx, provider, http.MethodGet, goCardlessAgreement+res.Agreement+"/", nil)
	if err != nil {
		return nil, err
	}

	respAgreement, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer respAgreement.Body.Close() //nolint: errcheck

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, model.ErrRateLimited
	}
	if respAgreement.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get agreement: " + respAgreement.Status)
	}

	var resAgreement agreementResponse
	if err := json.NewDecoder(respAgreement.Body).Decode(&resAgreement); err != nil {
		return nil, err
	}

	renewConsentAt := resAgreement.Accepted.AddDate(0, 0, resAgreement.AccessValidForDays)

	return &model.Connection{
		ProviderUserID:       providerUser.ID,
		ProviderConnectionID: connectionID,
		ConnectorID:          connector.ID,
		Status:               convertStatus(res.Status),
		RenewConsentBefore:   &renewConsentAt,
		ErrorMessage:         getErrorMessage(res.Status),
		LastSuccessfulSync:   nil,
	}, nil
}

func convertStatus(status string) model.ConnectionStatus {
	switch status {
	case "EX":
		return model.ConnectionStatusSuspended
	case "CR":
	case "GC":
	case "UA":
	case "SA":
	case "GA":
		return model.ConnectionStatusSyncInProgress
	case "LN":
		return model.ConnectionStatusSynchronized
	}
	return model.ConnectionStatusDeleted
}

func getErrorMessage(status string) *string {
	var msg string
	if status == "EX" {
		msg = "Access to accounts has expired as set in End User Agreement"
	} else if status == "RJ" {
		msg = "Either SSN verification has failed or end-user has entered incorrect credentials"
	}
	return &msg
}

type requisitionResponse struct {
	Id                string    `json:"id"`
	Created           time.Time `json:"created"`
	Redirect          string    `json:"redirect"`
	Status            string    `json:"status"`
	InstitutionId     string    `json:"institution_id"`
	Agreement         string    `json:"agreement"`
	Reference         string    `json:"reference"`
	Accounts          []string  `json:"accounts"`
	UserLanguage      string    `json:"user_language"`
	Link              string    `json:"link"`
	Ssn               string    `json:"ssn"`
	AccountSelection  bool      `json:"account_selection"`
	RedirectImmediate bool      `json:"redirect_immediate"`
}

type agreementResponse struct {
	Id                 string    `json:"id"`
	Created            time.Time `json:"created"`
	InstitutionId      string    `json:"institution_id"`
	MaxHistoricalDays  int       `json:"max_historical_days"`
	AccessValidForDays int       `json:"access_valid_for_days"`
	AccessScope        []string  `json:"access_scope"`
	Accepted           time.Time `json:"accepted"`
}
