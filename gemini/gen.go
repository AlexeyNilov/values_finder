package gemini

import (
	"context"
	"log"
	"os"

	"github.com/AlexeyNilov/values_finder/core"
	"google.golang.org/genai"
)

type Client struct {
	Model string
}

func (c Client) GenText(prompt string) string {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_GENAI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		c.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text()
}

func (c Client) GenerateOptions(history []core.Choice) ([]string, error) {
	return nil, nil
}

func (c Client) GenerateFinalValues(history []core.Choice) (core.RankedValues, error) {
	return nil, nil
}
