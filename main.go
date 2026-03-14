package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"finance-notification/client"
	"finance-notification/model"
)

func handler(ctx context.Context) ([]model.QuoteOutput, error) {
	token := os.Getenv("BRAPI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("BRAPI_TOKEN não definido")
	}

	brapiClient := client.NewBrapiClient(token)

	stocks, err := brapiClient.FetchAllTickers()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar lista de tickers: %w", err)
	}

	log.Printf("Total de ações encontradas: %d", len(stocks))

	var results []model.QuoteOutput

	for _, s := range stocks {
		quote, err := brapiClient.FetchQuote(s.Stock)
		if err != nil {
			log.Printf("Erro ao buscar cotação de %s: %v", s.Stock, err)
			continue
		}
		if quote == nil {
			continue
		}

		results = append(results, model.ToQuoteOutput(*quote))
	}

	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar resultado: %w", err)
	}
	fmt.Println(string(data))

	return results, nil
}

func main() {
	lambda.Start(handler)
}
