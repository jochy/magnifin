package gocardless

import (
	"context"
	"magnifin/internal/app/model"
	"strconv"
)

func (g *GoCardless) CreateProviderUser(_ context.Context, provider *model.Provider, user *model.User) (*model.ProviderUser, error) {
	// No concept of users in GoCardless, so fake it
	return &model.ProviderUser{
		ProviderID:     provider.ID,
		UserID:         user.ID,
		ProviderUserID: strconv.Itoa(int(user.ID)),
	}, nil
}
