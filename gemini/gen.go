package gemini

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/AlexeyNilov/values_finder/core"
	"github.com/AlexeyNilov/values_finder/util"
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

type Options struct {
	Options []string `json:"options"`
}

// Function to parse JSON and return []string
func ParseOptions(input string) []string {
	// Remove Markdown code fences
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "```json")
	input = strings.TrimSuffix(input, "```")
	input = strings.TrimSpace(input)

	var opts Options
	err := json.Unmarshal([]byte(input), &opts)
	if err != nil {
		log.Fatal(err)
	}
	return opts.Options
}

func (c Client) GenerateOptions(history []core.Choice) ([]string, error) {
	template := util.ReadTemplate("doc/options_prompt.md")
	// TODO Transform history into string form
	data := struct {
		Data string
	}{
		Data: "[]",
	}
	prompt := util.ParseTemplate(template, data)
	result := c.GenText(prompt)
	options := ParseOptions(result)
	return options, nil
}

func (c Client) GenerateFinalValues(history []core.Choice) (core.RankedValues, error) {
	return nil, nil
}
