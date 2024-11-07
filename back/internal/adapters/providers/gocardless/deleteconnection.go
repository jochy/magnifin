package gocardless

import (
	"context"
	"errors"
	"fmt"
	"magnifin/internal/app/model"
	"net/http"
	"strconv"
)

func (g *GoCardless) DeleteConnection(
	ctx context.Context,
	provider *model.Provider,
	_ *model.ProviderUser,
	connection *model.Connection,
) error {
	req, err := g.newRequest(ctx, provider, http.MethodDelete, goCardlessRequisition+connection.ProviderConnectionID+"/", nil)
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("unable to send delete requisition request: %w", err)
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode >= 300 {
		return errors.New("unexpected status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
