package core

import "time"

// Choice represents a single user choice during a session
type Choice struct {
	QuestionText string   // Full text of the question prompt
	Options      []string // List of options presented
	Selected     int      // Index of user-selected option
}

// RankedValue represents a single value in the final ranked list
type RankedValue struct {
	Name        string
	Description string
}

// RankedValues is a slice of RankedValue representing the final ranked list
type RankedValues []RankedValue

// SessionData contains all data for a complete session
type SessionData struct {
	Timestamp    time.Time
	Choices      []Choice
	FinalRanking RankedValues
}
