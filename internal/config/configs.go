package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
	Testing     Environment = "testing"
)

type Config struct {
	Environment Environment
	Port        string
	LogLevel    string
	LogFile     string
	Debug       bool
}

func LoadConfig() (*Config, error) {
	// Load environment-specific .env file
	env := getEnvironment()
	envFile := fmt.Sprintf(".env.%s", env)

	// Try to load environment-specific file first, then fallback to .env
	if err := godotenv.Load(envFile); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("No environment file found (%s or .env)", envFile)
		}
	}

	port := ":3000"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	config := &Config{
		Environment: env,
		Port:        port,
		LogLevel:    getEnvWithDefault("LOG_LEVEL", getDefaultLogLevel(env)),
		LogFile:     getEnvWithDefault("LOG_FILE", getDefaultLogFile(env)),
		Debug:       getEnvWithDefault("DEBUG", getDefaultDebug(env)) == "true",
	}

	log.Printf("Loaded configuration for environment: %s", config.Environment)
	return config, nil
}

func getEnvironment() Environment {
	env := strings.ToLower(os.Getenv("APP_ENV"))
	switch env {
	case "production", "prod":
		return Production
	case "testing", "test":
		return Testing
	default:
		return Development
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDefaultLogLevel(env Environment) string {
	switch env {
	case Production:
		return "info"
	case Testing:
		return "error"
	default:
		return "debug"
	}
}

func getDefaultLogFile(env Environment) string {
	switch env {
	case Production:
		return "storage/logs/production.log"
	case Testing:
		return "storage/logs/testing.log"
	default:
		return "storage/logs/development.log"
	}
}

func getDefaultDebug(env Environment) string {
	switch env {
	case Production:
		return "false"
	default:
		return "true"
	}
}
