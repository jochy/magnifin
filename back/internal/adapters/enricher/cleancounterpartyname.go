package enricher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"log/slog"
	"magnifin/internal/app/transactions"
	"net/http"
	"strings"
)

type OpenAICounterpartyNameResponse struct {
	PaymentMethod    *string `json:"payment_method" jsonschema_description:"The payment method used in the transaction"`
	CounterpartyName *string `json:"counterparty_name" jsonschema_description:"The name of the counterparty in the transaction" enum:"SEPA,CARD,CASH,CHECK,DIRECT DEBIT,PAYPAL,SWIFT,OTHER"`
}

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var CounterpartyNameResponseSchema = GenerateSchema[OpenAICounterpartyNameResponse]()

func (e *Enricher) CleanCounterpartyName(ctx context.Context, name *string, userCounterparties []string) (*transactions.CounterpartyEnrichmentResult, error) {
	if e.openAIClient == nil {
		slog.Debug("CleanCounterpartyName: openAI client is not initialized, skipping the counterparty name enrichment")
		return nil, nil
	}

	if name == nil {
		slog.Debug("CleanCounterpartyName: counterparty name is nil, skipping the counterparty name enrichment")
		return nil, nil
	}

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("counterparty"),
		Description: openai.F("Counterparty transaction data"),
		Schema:      openai.F(CounterpartyNameResponseSchema),
		Strict:      openai.Bool(true),
	}

	chatCompletion, err := e.openAIClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("you are a bot, used by another program to extract data from a transaction reference. Reply only with a json and without any formatting. Your data must be as clean as possible, so remove any duplicated word, space. Remove any identifier (i.e. SCA00043K134, SCR51684, ICS.05651, RUM.51684 are forbidden). No dates. Return only a JSON with `payment_method` and `counterparty_name` entries. If you are not sure, return an empty json."),
			openai.SystemMessage("payment_method expected value: SEPA, CARD, CASH, CHECK, DIRECT DEBIT, PAYPAL, SWIFT, OTHER (if you don't know the payment method). Custom mapping: CB->CARD, CARTE->CARD, VIREMENT->SEPA"),
			openai.SystemMessage("counterparty_name expected value: the name of the counterparty in the transaction (which is the company, or the name of the debtor, usually, the first name after the payment method). **FORBIDDEN** word in this field: CONFRERE, D/O, RECU, PRLV, please respect this rule!"),
			openai.SystemMessage("To help you, here are all known counterparties: " + strings.Join(userCounterparties, ", ") + ". Use them **only if you find a match**."),
			openai.UserMessage(*name),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Model:       openai.F(e.aiModel),
		TopP:        openai.Float(0.1),
		Temperature: openai.Float(0.2),
	}, option.WithMiddleware(ollamaMiddleware()))
	if err != nil {
		return nil, fmt.Errorf("cleanCounterpartyName: failed to get the completion: %w", err)
	}

	counterparty := OpenAICounterpartyNameResponse{}
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &counterparty)

	aiCounterparty := ""
	if counterparty.CounterpartyName != nil {
		aiCounterparty = strings.TrimSpace(*counterparty.CounterpartyName)
		counterparty.CounterpartyName = &aiCounterparty
	}
	rankMatch := fuzzy.RankMatchNormalizedFold(aiCounterparty, *name)
	if rankMatch == -1 {
		slog.Debug("CleanCounterpartyName: counterparty name is not a match, ignoring the counterparty name enrichment")
		return &transactions.CounterpartyEnrichmentResult{
			CounterpartyName: nil,
			Method:           counterparty.PaymentMethod,
		}, nil
	}

	return &transactions.CounterpartyEnrichmentResult{
		CounterpartyName: counterparty.CounterpartyName,
		Method:           counterparty.PaymentMethod,
	}, nil
}

func ollamaMiddleware() func(request *http.Request, next option.MiddlewareNext) (*http.Response, error) {
	return func(request *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		// Needed if use ollama
		if strings.Contains(request.URL.Path, "/v1") || strings.Contains(request.URL.Path, "api.openai.com") {
			return next(request)
		}

		request.URL.Path = "/v1" + request.URL.Path
		return next(request)
	}
}
