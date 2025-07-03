package utils

import (
	"fmt"
	"log"
	"os"
)

func SetupLogging() error {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("storage/logs", 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Open log file
	logFile, err := os.OpenFile("storage/logs/develop.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Set log output to file
	log.SetOutput(logFile)

	return nil
}
