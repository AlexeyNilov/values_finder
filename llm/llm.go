package llm

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/AlexeyNilov/values_finder/core"
	"github.com/AlexeyNilov/values_finder/gemini"
	"github.com/AlexeyNilov/values_finder/util"
)

// Client defines the interface for LLM operations
type Client interface {
	GenerateOptions(history []core.Choice) ([]string, error)
	GenerateFinalValues(history []core.Choice) (core.RankedValues, error)
}

type LLMClient struct {
	gemini.Client
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

func ParseRankedValues(input string) core.RankedValues {
    // Remove Markdown code fences if present
    input = strings.TrimSpace(input)
    input = strings.TrimPrefix(input, "```json")
    input = strings.TrimSuffix(input, "```")
    input = strings.TrimSpace(input)

    var values core.RankedValues
    err := json.Unmarshal([]byte(input), &values)
    if err != nil {
        log.Fatal(err)
    }

    return values
}


func ExtractPreviousOptions(history []core.Choice) []string {
	results := make([]string, 0, len(history))

	for i, choice := range history {
		if choice.Selected < 0 || choice.Selected >= len(choice.Options) {
			log.Fatalf("invalid Selected index at history[%d]: %d", i, choice.Selected)
		}
		// TODO add question "option 1 or options 2? choice"
		// record := choice.Options[0] + " or " + choice.Options[0] + "? User choice: " + choice.Options[choice.Selected]
		results = append(results, choice.Options[choice.Selected])
	}

	return results
}

func formatAsJSONString(values []string) string {
	data, err := json.MarshalIndent(values, "  ", "  ")
	if err != nil {
		log.Fatalf("failed to marshal values: %v", err)
	}

	return fmt.Sprintf("```json\n%s\n```", string(data))
}

func (c LLMClient) GenerateOptions(history []core.Choice) ([]string, error) {
	template := util.ReadTemplate("doc/options_prompt.md")
	prevOptions := ExtractPreviousOptions(history)
	data := struct {
		Data string
	}{
		Data: formatAsJSONString(prevOptions),
	}
	prompt := util.ParseTemplate(template, data)
	result := c.GenText(prompt)
	options := ParseOptions(result)
	return options, nil
}

func (c LLMClient) GenerateFinalValues(history []core.Choice) (core.RankedValues, error) {
	template := util.ReadTemplate("doc/rank_prompt.md")
	prevOptions := ExtractPreviousOptions(history)
	data := struct {
		Data string
	}{
		Data: formatAsJSONString(prevOptions),
	}
	prompt := util.ParseTemplate(template, data)
	result := c.GenText(prompt)
	fmt.Println(result)
	return ParseRankedValues(result), nil
}
