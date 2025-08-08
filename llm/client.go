package llm

import (
	"fmt"

	"github.com/AlexeyNilov/values_finder/core"
)

// Client defines the interface for LLM operations
type Client interface {
	GenerateOptions(history []core.Choice) ([]string, error)
	GenerateFinalValues(history []core.Choice) (core.RankedValues, error)
}

// MockClient implements the Client interface for testing
type MockClient struct {
	ShouldFail bool
}

// GenerateOptions returns hardcoded options for testing
func (m *MockClient) GenerateOptions(history []core.Choice) ([]string, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("mock error: failed to generate options")
	}

	// Return hardcoded options for testing
	return []string{"Being creative", "Being disciplined"}, nil
}

// GenerateFinalValues returns hardcoded ranked values for testing
func (m *MockClient) GenerateFinalValues(history []core.Choice) (core.RankedValues, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("mock error: failed to generate final values")
	}

	// Return hardcoded ranked values for testing
	return core.RankedValues{
		{
			Name:        "Creativity",
			Description: "You value expressing yourself and thinking outside the box. It's a core part of who you are.",
		},
		{
			Name:        "Discipline",
			Description: "You appreciate structure and the power of consistency to achieve long-term goals.",
		},
		{
			Name:        "Growth",
			Description: "You believe in continuous learning and personal development as a path to fulfillment.",
		},
	}, nil
}
