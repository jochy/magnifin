package gocardless

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"magnifin/internal/app/model"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	goCardlessURL = "https://bankaccountdata.gocardless.com"
)

type gocardlessAccessToken struct {
	Access         string `json:"access"`
	AccessExpires  int    `json:"access_expires"`
	Refresh        string `json:"refresh"`
	RefreshExpires int    `json:"refresh_expires"`

	IssuedAt *time.Time
}

type GoCardless struct {
	token     *gocardlessAccessToken
	mu        sync.Mutex
	publicURL string
}

func NewGoCardless(publicURL string) *GoCardless {
	return &GoCardless{
		token:     nil,
		mu:        sync.Mutex{},
		publicURL: publicURL,
	}
}

func (g *GoCardless) Name() string {
	return "GoCardless"
}

func (g *GoCardless) LoadConfig() *model.Provider {
	var accessKey *string
	var secret *string

	if os.Getenv("GOCARDLESS_SECRET_ID") != "" {
		tmp := os.Getenv("GOCARDLESS_SECRET_ID")
		accessKey = &tmp
	}
	if os.Getenv("GOCARDLESS_SECRET_KEY") != "" {
		tmp := os.Getenv("GOCARDLESS_SECRET_KEY")
		secret = &tmp
	}

	return &model.Provider{
		Name:      g.Name(),
		AccessKey: accessKey,
		Secret:    secret,
	}
}

func (g *GoCardless) ValidateConfiguration(provider *model.Provider) error {
	if provider.AccessKey == nil || provider.Secret == nil {
		slog.Error("Access key and secret are required")
		return errors.New("access key and secret are required")
	}

	slog.Debug("GoCardless provider configuration is valid")
	return nil
}

func (g *GoCardless) newRequest(
	ctx context.Context,
	provider *model.Provider,
	method string,
	url string,
	body []byte,
) (*http.Request, error) {
	err := g.updateTokenIfNeeded(ctx, provider)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, goCardlessURL+url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+g.token.Access)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}
