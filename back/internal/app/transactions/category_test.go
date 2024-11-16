package transactions

import (
	"magnifin/internal/app/model"
	"testing"
)

func Test_evaluateRule(t *testing.T) {
	type args struct {
		rule        model.CategoryRule
		transaction model.Transaction
		enrichment  model.TransactionEnrichment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Rule matches transaction and enrichment",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{"100.00", "USD", "123", "outgoing", "Amazon", "Credit Card", "Order123"},
				},
				transaction: model.Transaction{
					Amount:    100.00,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: true,
		},
		{
			name: "Rule does not match transaction and enrichment",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{"200.00", "EUR", "456", "incoming", "Ebay", "PayPal", "Order456"},
				},
				transaction: model.Transaction{
					Amount:    100.00,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: false,
		},
		{
			name: "Rule does not match transaction and enrichment",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{"100.00", "USD", "123", "outgoing", "Amazon", "Credit Card", "Order123"},
				},
				transaction: model.Transaction{
					Amount:    100.01,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: false,
		},
		{
			name: "Rule matches transaction and enrichment with large reference and ignore case",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{"100.00", "USD", "123", "outgoing", "Amazon", "Credit Card", "Order123", "from", "paris"},
				},
				transaction: model.Transaction{
					Amount:    100.00,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123 from Paris"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: true,
		},
		{
			name: "Rule partially matches transaction and enrichment",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{"100.00", "USD", "123", "outgoing", "Amazon"},
				},
				transaction: model.Transaction{
					Amount:    100.00,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: true,
		},
		{
			name: "Empty rule",
			args: args{
				rule: model.CategoryRule{
					Rule: []string{},
				},
				transaction: model.Transaction{
					Amount:    100.00,
					Currency:  "USD",
					AccountID: 123,
					Direction: "outgoing",
					Reference: ptr("Order123"),
				},
				enrichment: model.TransactionEnrichment{
					CounterpartyName: ptr("Amazon"),
					Method:           ptr("Credit Card"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateRule(tt.args.rule, transactionKeywords(&tt.args.transaction, &tt.args.enrichment)); got != tt.want {
				t.Errorf("evaluateRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ptr(s string) *string {
	return &s
}
