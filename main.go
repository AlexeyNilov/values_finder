package main

import (
	"log"

	"github.com/AlexeyNilov/values_finder/config"
)

func main() {
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Configuration loaded: %d rounds.", config.Rounds)
}
