package main

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/genai"
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
	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{
			APIKey: cfg.APIKey,
		},
	)
	if err != nil {
		panic(err)
	}

	for item, err := range client.Models.All(ctx) {
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", item.Name)
	}

	model := "gemini-2.0-flash-lite-001"
	prompt := "Introduce yourself in 1-2 sentences."

	t, err := client.Models.CountTokens(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("total token: %d", t.TotalTokens))

	resp, err := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Text())

}
