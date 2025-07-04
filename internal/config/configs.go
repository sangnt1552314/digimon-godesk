package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := ":3000"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	} else {
		log.Println("Using default port 3000")
	}

	return &Config{
		Port: port,
	}, nil
}
