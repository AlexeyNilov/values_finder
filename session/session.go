package session

import (
	"fmt"
	"os"
	"time"

	"github.com/AlexeyNilov/values_finder/core"
)

// Manager handles session data and file logging
type Manager struct {
	sessionData core.SessionData
	file        *os.File
}

// NewManager creates a new session manager with a timestamped log file
func NewManager() (*Manager, error) {
	// Create session data with current timestamp
	sessionData := core.SessionData{
		Timestamp: time.Now(),
		Choices:   []core.Choice{},
	}

	// Create timestamped filename
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("values_session_%s.txt", timestamp)

	// Create the log file
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create session file: %v", err)
	}

	// Write session header
	header := fmt.Sprintf("Values Discovery Session - %s\n", sessionData.Timestamp.Format("2006-01-02 15:04:05"))
	header += "=====================================\n\n"
	_, err = file.WriteString(header)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to write session header: %v", err)
	}

	return &Manager{
		sessionData: sessionData,
		file:        file,
	}, nil
}

// AddChoice appends a choice to the session history and logs it to the file
func (m *Manager) AddChoice(choice core.Choice) error {
	// Add to session data
	m.sessionData.Choices = append(m.sessionData.Choices, choice)

	// Log to file
	logEntry := fmt.Sprintf("Question: %s\n", choice.QuestionText)
	for i, option := range choice.Options {
		logEntry += fmt.Sprintf("  %d) %s\n", i+1, option)
	}
	logEntry += fmt.Sprintf("Selected: %d) %s\n\n", choice.Selected+1, choice.Options[choice.Selected])

	_, err := m.file.WriteString(logEntry)
	if err != nil {
		return fmt.Errorf("failed to write choice to log file: %v", err)
	}

	return nil
}

// LogFinalValues writes the final ranked values to the log file
func (m *Manager) LogFinalValues(values core.RankedValues) error {
	// Store in session data
	m.sessionData.FinalRanking = values

	// Log to file
	logEntry := "\nFinal Values:\n"
	logEntry += "============\n\n"

	for i, value := range values {
		logEntry += fmt.Sprintf("%d. %s\n", i+1, value.Name)
		logEntry += fmt.Sprintf("   %s\n\n", value.Description)
	}

	_, err := m.file.WriteString(logEntry)
	if err != nil {
		return fmt.Errorf("failed to write final values to log file: %v", err)
	}

	return nil
}

// GetHistory returns the choice history
func (m *Manager) GetHistory() []core.Choice {
	return m.sessionData.Choices
}

// Close closes the session file
func (m *Manager) Close() error {
	if m.file != nil {
		return m.file.Close()
	}
	return nil
}
