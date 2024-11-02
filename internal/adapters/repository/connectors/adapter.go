package connectors

import (
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

func toDomain(connector database.Connector) *model.Connector {
	return &model.Connector{
		ID:                  connector.ID,
		ProviderID:          connector.ProviderID,
		ProviderConnectorID: connector.ProviderConnectorID,
		Name:                connector.Name,
		LogoURL:             connector.LogoUrl.String,
	}
}
