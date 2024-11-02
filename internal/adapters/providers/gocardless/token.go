package gocardless

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"
	"net/http"
)

const (
	goCardlessNewToken = "/api/v2/token/new" //nolint: gosec
)

func (g *GoCardless) updateTokenIfNeeded(ctx context.Context, provider *model.Provider) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	switch {
	case g.token == nil:
		return g.generateNewToken(ctx, provider)
	case g.token.AccessExpires < 10 && g.token.RefreshExpires > 10:
		err := g.updateToken(ctx)
		if err != nil {
			slog.Warn(fmt.Sprintf("Failed to update token: %s, will request a brand new one...", err.Error()))
			return g.generateNewToken(ctx, provider)
		}
	case g.token.AccessExpires < 10:
		return g.generateNewToken(ctx, provider)
	}

	return nil
}

func (g *GoCardless) generateNewToken(ctx context.Context, provider *model.Provider) error {
	slog.Info("Generating new token for GoCardless provider")

	type tokenRequest struct {
		SecretID  string `json:"secret_id"`
		SecretKey string `json:"secret_key"`
	}

	tokenReq := tokenRequest{SecretID: *provider.AccessKey, SecretKey: *provider.Secret}

	b, err := json.Marshal(tokenReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, goCardlessURL+goCardlessNewToken, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to generate new token http status code is %s" + resp.Status)
	}

	var token gocardlessAccessToken
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return err
	}

	g.token = &token

	return nil
}

func (g *GoCardless) updateToken(ctx context.Context) error {
	slog.Info("Updating token for GoCardless provider")

	type tokenRefreshRequest struct {
		Refresh string `json:"refresh"`
	}

	tokenReq := tokenRefreshRequest{Refresh: g.token.Refresh}

	b, err := json.Marshal(tokenReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, goCardlessURL+goCardlessNewToken, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to generate new token http status code is %s" + resp.Status)
	}

	var token gocardlessAccessToken
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return err
	}

	g.token.Access = token.Access
	g.token.AccessExpires = token.AccessExpires

	return nil
}
