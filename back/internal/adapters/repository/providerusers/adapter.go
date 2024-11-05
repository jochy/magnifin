package providerusers

import (
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

func toDomain(providerUser *database.ProviderUser) *model.ProviderUser {
	return &model.ProviderUser{
		ID:             providerUser.ID,
		ProviderID:     providerUser.ProviderID,
		UserID:         providerUser.UserID,
		ProviderUserID: providerUser.ProviderUserID,
	}
}
