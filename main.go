package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Setup logging
	if errLog := os.MkdirAll("storage/logs", 0755); errLog != nil {
		panic(fmt.Errorf("failed to create logs directory: %w", errLog))
	}

	logFile, errLogFile := os.OpenFile("storage/logs/develop.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errLogFile != nil {
		panic(errLogFile)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	//Setup environment variables
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := ":3000"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	} else {
		log.Println("Using default port 3000")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Digimon GoDesk is running!")
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
