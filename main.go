package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlexeyNilov/values_finder/config"
	"github.com/AlexeyNilov/values_finder/core"
	"github.com/AlexeyNilov/values_finder/gemini"
	"github.com/AlexeyNilov/values_finder/session"
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Initialize session manager
	manager, err := session.NewManager()
	if err != nil {
		log.Fatalf("Failed to create session manager: %v", err)
	}
	defer manager.Close()

	// 3. Initialize LLM client (using mock for now)
	// client := &llm.MockClient{ShouldFail: false}
	client := &gemini.Client{Model: "gemini-2.0-flash-lite"}

	// 4. Main interaction loop
	for round := 1; round <= cfg.Rounds; round++ {
		// Display progress
		fmt.Printf("Round %d of %d\n", round, cfg.Rounds)

		// Generate options from LLM
		options, err := client.GenerateOptions(manager.GetHistory())
		if err != nil {
			log.Fatalf("Failed to generate options: %v", err)
		}

		// Display question and options
		fmt.Println("Which feels more important to you right now:")
		for i, option := range options {
			fmt.Printf("%d) %s\n", i+1, option)
		}

		// Get user input
		var input string
		fmt.Print("Enter your choice (1 or 2): ")
		_, _ = fmt.Scanln(&input)

		// Parse user input (convert from 1-based to 0-based index)
		selected, err := strconv.Atoi(input)
		if err != nil {
			log.Printf("Invalid input, defaulting to option 1")
			selected = 1
		}

		// Ensure valid range
		if selected < 1 || selected > len(options) {
			selected = 1
		}

		// Create choice struct
		choice := core.Choice{
			QuestionText: "Which feels more important to you right now:",
			Options:      options,
			Selected:     selected - 1, // Convert to 0-based index
		}

		// Add choice to session
		err = manager.AddChoice(choice)
		if err != nil {
			log.Printf("Failed to log choice: %v", err)
		}

		fmt.Println() // Add spacing between rounds
	}

	// 5. Generate final values
	fmt.Println("Generating your values...")

	finalValues, err := client.GenerateFinalValues(manager.GetHistory())
	if err != nil {
		log.Fatalf("Failed to generate final values: %v", err)
	}

	// 6. Display final values
	fmt.Println("\nHere's what seems most important to you:")
	for i, value := range finalValues {
		fmt.Printf("%d. %s\n", i+1, value.Name)
		fmt.Printf("   %s\n\n", value.Description)
	}

	// 7. Log final values to session file
	err = manager.LogFinalValues(finalValues)
	if err != nil {
		log.Printf("Failed to log final values: %v", err)
	}
}
