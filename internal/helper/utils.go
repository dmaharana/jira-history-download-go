package helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"jira-history-download/internal/jira"
)

func WriteToCSV(items []jira.HistoryItem, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Issue Key", "Author", "Created Date", "Field", "Old Value", "New Value"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	// Write data
	for _, item := range items {
		row := []string{
			item.IssueKey,
			item.Author,
			item.CreatedDate,
			item.Field,
			item.OldValue,
			item.NewValue,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %v", err)
		}
	}

	return nil
}

// generate a timestamped filename
func GenerateFilename(basefilename string) string {
	return fmt.Sprintf("%s_%s.csv", basefilename, time.Now().Format("2006-01-02_150405"))
}

// check if folder exists, if not create it
func EnsureFolderExists(foldername string) error {
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		if err := os.Mkdir(foldername, 0755); err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
	}
	return nil
}