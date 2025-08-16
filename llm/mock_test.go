package llm

import (
	"testing"

	"github.com/AlexeyNilov/values_finder/core"
)

func TestMockClient_GenerateOptions(t *testing.T) {
	client := &MockClient{ShouldFail: false}

	options, err := client.GenerateOptions([]core.Choice{})
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(options))
	}

	expectedOptions := []string{"Being creative", "Being disciplined"}
	for i, option := range options {
		if option != expectedOptions[i] {
			t.Errorf("Expected option %d to be '%s', got '%s'", i, expectedOptions[i], option)
		}
	}
}

func TestMockClient_GenerateOptions_Error(t *testing.T) {
	client := &MockClient{ShouldFail: true}

	_, err := client.GenerateOptions([]core.Choice{})
	if err == nil {
		t.Error("Expected error when ShouldFail is true")
	}

	if err.Error() != "mock error: failed to generate options" {
		t.Errorf("Expected specific error message, got: %v", err)
	}
}

func TestMockClient_GenerateFinalValues(t *testing.T) {
	client := &MockClient{ShouldFail: false}

	values, err := client.GenerateFinalValues([]core.Choice{})
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Check first value
	if values[0].Name != "Creativity" {
		t.Errorf("Expected first value name to be 'Creativity', got '%s'", values[0].Name)
	}

	if values[0].Description == "" {
		t.Error("Expected first value to have a description")
	}

	// Check second value
	if values[1].Name != "Discipline" {
		t.Errorf("Expected second value name to be 'Discipline', got '%s'", values[1].Name)
	}

	// Check third value
	if values[2].Name != "Growth" {
		t.Errorf("Expected third value name to be 'Growth', got '%s'", values[2].Name)
	}
}

func TestMockClient_GenerateFinalValues_Error(t *testing.T) {
	client := &MockClient{ShouldFail: true}

	_, err := client.GenerateFinalValues([]core.Choice{})
	if err == nil {
		t.Error("Expected error when ShouldFail is true")
	}

	if err.Error() != "mock error: failed to generate final values" {
		t.Errorf("Expected specific error message, got: %v", err)
	}
}
