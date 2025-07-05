package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sangnt1552314/digimon-godesk/internal/config"
)

func SetupLogging(cfg *config.Config) error {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("storage/logs", 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	var output io.Writer

	if cfg.Environment == config.Production {
		// In production, write only to file
		logFile, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		output = logFile
	} else {
		// In development, write to both file and stdout
		logFile, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		output = io.MultiWriter(os.Stdout, logFile)
	}

	// Set log output
	log.SetOutput(output)

	// Set log flags based on environment
	if cfg.Environment == config.Production {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	}

	// Log level filtering could be implemented here
	logLevel := strings.ToLower(cfg.LogLevel)
	log.Printf("Logging initialized with level: %s, file: %s", logLevel, cfg.LogFile)

	return nil
}
