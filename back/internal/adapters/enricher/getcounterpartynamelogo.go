package enricher

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

const brandApiUrl = "https://api.brandfetch.io/v2/search/%s?limit=5"

func (e *Enricher) GetCounterpartyNameLogoURL(ctx context.Context, name *string) (*string, error) {
	if name == nil || *name == "" {
		slog.Debug("GetCounterpartyNameLogoURL: name is nil, skipping")
		return nil, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(brandApiUrl, *name), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var searchResponse []brandSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	if len(searchResponse) == 0 {
		return nil, nil
	}

	return &searchResponse[0].Icon, nil
}

type brandSearchResponse struct {
	BrandId  string `json:"brandId"`
	Claimed  bool   `json:"claimed"`
	Domain   string `json:"domain"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Verified bool   `json:"verified"`
}
