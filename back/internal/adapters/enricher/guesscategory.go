package enricher

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func (e *Enricher) GuessCategory(ctx context.Context, keywords []string, categories []string) (*string, error) {
	if e.openAIClient == nil {
		slog.Debug("openAI client is not initialized, skipping the guess category")
		return nil, nil
	}

	if len(categories) == 0 {
		slog.Debug("no categories provided, skipping the guess category")
		return nil, nil
	}

	if len(keywords) == 0 {
		slog.Debug("no keywords provided, skipping the guess category")
		return nil, nil
	}

	chatCompletion, err := e.openAIClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("you are a bot, used by another program to guess a transaction category, based on keywords. Reply only with the category name, nothing else. If you are not sure, return ###. You have to chose one category among the following: " + strings.Join(categories, ", ")),
			openai.UserMessage(strings.Join(keywords, " ")),
		}),
		Model:       openai.F(e.aiModel),
		TopP:        openai.Float(0.1),
		Temperature: openai.Float(0.2),
	}, option.WithMiddleware(ollamaMiddleware())) //nolint: bodyclose
	if err != nil {
		return nil, fmt.Errorf("cleanCounterpartyName: failed to get the completion: %w", err)
	}

	return &chatCompletion.Choices[0].Message.Content, nil
}
