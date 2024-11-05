package gocardless

import (
	"context"
	"encoding/json"
	"errors"
	"magnifin/internal/app/model"
	"net/http"
	"net/url"
	"strconv"
)

const (
	goCardlessRequisition = "/api/v2/requisitions/"
)

func (g *GoCardless) Connect(
	ctx context.Context,
	provider *model.Provider,
	_ *model.ProviderUser,
	connector *model.Connector,
	params *model.ConnectParams,
) (*model.ConnectInstruction, error) {
	// Could be handled better, but it does the job...
	c, err := g.sendRequisitionRequest(ctx, provider, params, connector, true)
	if err != nil {
		return g.sendRequisitionRequest(ctx, provider, params, connector, false)
	}
	return c, nil
}

func (g *GoCardless) sendRequisitionRequest(
	ctx context.Context,
	provider *model.Provider,
	params *model.ConnectParams,
	connector *model.Connector,
	withAccountSelection bool,
) (*model.ConnectInstruction, error) {
	u, err := url.Parse(g.publicURL + "/v1/providers/gocardless/callback")
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Add("s", params.SuccessURL)
	queryParams.Add("e", params.ErrorURL)
	queryParams.Add("c", strconv.Itoa(int(connector.ID)))
	queryParams.Add("sid", params.SID.String())
	u.RawQuery = queryParams.Encode()

	reqBody := requisitionRequest{
		Redirect:         u.String(),
		InstitutionId:    connector.ProviderConnectorID,
		AccountSelection: withAccountSelection,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := g.newRequest(ctx, provider, http.MethodPost, goCardlessRequisition, body)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("failed to create redirection link: " + resp.Status)
	}

	var res requisitionResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &model.ConnectInstruction{
		ID:          res.Id,
		RedirectURL: res.Link,
	}, nil
}

type requisitionRequest struct {
	Redirect         string `json:"redirect"`
	InstitutionId    string `json:"institution_id"`
	AccountSelection bool   `json:"account_selection"`
}
