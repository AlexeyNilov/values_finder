package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test valid config
	validConfig := `rounds: 20
options_per_question: 2`

	tmpFile, err := os.CreateTemp("", "config_test_*.yml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(validConfig); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	config, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("LoadConfig failed with valid config: %v", err)
	}

	if config.Rounds != 20 {
		t.Errorf("Expected Rounds to be 20, got %d", config.Rounds)
	}

	if config.OptionsPerQuestion != 2 {
		t.Errorf("Expected OptionsPerQuestion to be 2, got %d", config.OptionsPerQuestion)
	}

	// Test non-existent file
	_, err = LoadConfig("non_existent_file.yml")
	if err == nil {
		t.Error("Expected error when loading non-existent file, got nil")
	}
}
