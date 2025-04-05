package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/api/option"
)

func main() {
	var cfg struct {
		APIKey string `envconfig:"API_KEY" required:"true"`
	}

	err := envconfig.Process("GENAI", &cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.APIKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	modelName := "gemini-2.0-flash-lite-001"
	prompt := "Introduce yourself in 1-2 sentences."
  model := client.GenerativeModel(modelName)
	resp, err := model.GenerateContent(
    ctx, 
    genai.Text(prompt),
  )
	if err != nil {
		panic(err)
	}
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
