package enricher

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Enricher struct {
	openAIClient *openai.Client
	aiModel      string
}

func NewEnricher(openAIBaseUrl string, openAIKey string, aiModel string) *Enricher {
	var openAIClient *openai.Client
	if openAIKey != "" || openAIBaseUrl != "" {
		openAIClient = openai.NewClient(
			option.WithAPIKey(openAIKey),
			option.WithBaseURL(openAIBaseUrl),
		)
	}

	return &Enricher{
		openAIClient: openAIClient,
		aiModel:      aiModel,
	}
}
