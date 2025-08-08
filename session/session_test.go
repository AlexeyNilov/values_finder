package session

import (
	"os"
	"strings"
	"testing"

	"github.com/AlexeyNilov/values_finder/core"
)

func TestSessionFlow(t *testing.T) {
	// Create a new session manager
	manager, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create session manager: %v", err)
	}
	defer func() {
		if manager.file != nil {
			manager.file.Close()
			os.Remove(manager.file.Name())
		}
	}()

	// Verify the file exists
	if manager.file == nil {
		t.Fatal("Session file was not created")
	}

	// Check if the file name follows the expected pattern
	fileName := manager.file.Name()
	if !strings.Contains(fileName, "values_session_") {
		t.Errorf("Expected file name to contain 'values_session_', got: %s", fileName)
	}

	// Create a sample choice
	sampleChoice := core.Choice{
		QuestionText: "Which feels more important to you right now:",
		Options:      []string{"Being creative", "Being disciplined"},
		Selected:     0,
	}

	// Add the choice to the session
	err = manager.AddChoice(sampleChoice)
	if err != nil {
		t.Fatalf("Failed to add choice: %v", err)
	}

	// Read the content of the log file
	content, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	contentStr := string(content)

	// Assert that the file contains the question and selected option
	if !strings.Contains(contentStr, "Which feels more important to you right now:") {
		t.Error("Log file should contain the question text")
	}

	if !strings.Contains(contentStr, "Being creative") {
		t.Error("Log file should contain the selected option")
	}

	// Test GetHistory
	history := manager.GetHistory()
	if len(history) != 1 {
		t.Errorf("Expected 1 choice in history, got %d", len(history))
	}

	if history[0].QuestionText != sampleChoice.QuestionText {
		t.Error("History should contain the added choice")
	}

	// Test LogFinalValues
	finalValues := core.RankedValues{
		{Name: "Creativity", Description: "You value expressing yourself."},
		{Name: "Discipline", Description: "You appreciate structure and consistency."},
	}

	err = manager.LogFinalValues(finalValues)
	if err != nil {
		t.Fatalf("Failed to log final values: %v", err)
	}

	// Read the file again to check final values were logged
	content, err = os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read log file after final values: %v", err)
	}

	contentStr = string(content)
	if !strings.Contains(contentStr, "Creativity") || !strings.Contains(contentStr, "Discipline") {
		t.Error("Log file should contain the final ranked values")
	}
}
