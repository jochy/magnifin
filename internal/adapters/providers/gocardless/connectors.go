package gocardless

import (
	"context"
	"encoding/json"
	"errors"
	"magnifin/internal/app/model"
	"net/http"
)

const (
	goCardlessConnectors = "/api/v2/institutions"
)

func (g *GoCardless) ListConnectors(ctx context.Context, provider *model.Provider) ([]model.Connector, error) {
	req, err := g.newRequest(ctx, provider, http.MethodGet, goCardlessConnectors, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to list connectors http status code is %s" + resp.Status)
	}

	var connectors []connectorResponse
	err = json.NewDecoder(resp.Body).Decode(&connectors)
	if err != nil {
		return nil, err
	}

	result := make([]model.Connector, 0, len(connectors))
	for i, c := range connectors {
		result[i] = *c.toDomain(provider.ID)
	}

	return result, nil
}

type connectorResponse struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Bic                  string   `json:"bic"`
	TransactionTotalDays string   `json:"transaction_total_days"`
	Countries            []string `json:"countries"`
	Logo                 string   `json:"logo"`
}

func (c *connectorResponse) toDomain(providerID int32) *model.Connector {
	return &model.Connector{
		Name:                c.Name,
		ProviderConnectorID: c.Id,
		ProviderID:          providerID,
		LogoURL:             c.Logo,
	}
}
