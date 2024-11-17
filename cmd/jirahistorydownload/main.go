package main

import (
	"fmt"
	"log"
	"path/filepath"

	"jira-history-download/internal/config"
	"jira-history-download/internal/helper"
	"jira-history-download/internal/jira"
)

const (
	OutputDir = "output"
	OutputFile = "jira_history.csv"
)

func main() {
	// Load configuration
	configPath := filepath.Join("configs", "config.ini")
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create Jira client
	client, err := jira.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Jira client: %v", err)
	}

	// Get history for all issues
	allHistoryItems, err := client.GetHistoryForAllIssues(cfg.JQL)
	if err != nil {
		log.Fatalf("Error getting history for all issues: %v", err)
	}

	// Save history to CSV
	filename := helper.GenerateFilename(filepath.Join(OutputDir, OutputFile))
	err = saveToCSV(allHistoryItems, filename)
	if err != nil {
		log.Fatalf("Error saving history to CSV: %v", err)
	}
	
	fmt.Printf("Successfully exported history to %s\n", filename)
}

// save results to CSV
func saveToCSV(historyItems []jira.HistoryItem, filename string) error {
	// check if folder exists, if not create it
	if err := helper.EnsureFolderExists(OutputDir); err != nil {
		return err
	}
	
	// write to CSV
	return helper.WriteToCSV(historyItems, filename)
}
