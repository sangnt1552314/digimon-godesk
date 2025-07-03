package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/sangnt1552314/digimon-godesk/internal/utils"
)

func main() {
	// Setup logging
	if err := utils.SetupLogging(); err != nil {
		panic(err)
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := ":3000"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	} else {
		log.Println("Using default port 3000")
	}

	router := http.NewServeMux()

	//Handles static files
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
