package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/sangnt1552314/digimon-godesk/internal/models"
)

const (
	baseURL = "http://digi-api.com/api/v1/digimon"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the template
	tmplPath := "templates/index/index.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Fetch digimon data
	digimon, err := getDigimonByName("agumon") // Example name, can be replaced with a dynamic value
	if err != nil {
		http.Error(w, "Failed to fetch digimon", http.StatusInternalServerError)
		return
	}

	// Execute the template
	if err := tmpl.Execute(w, digimon); err != nil {
		log.Printf("Template execution error: %v", err)
		return
	}
}

func getDigimonByName(name string) (*models.DigimonDetail, error) {
	url := fmt.Sprintf("%s/%s", baseURL, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch digimon by name: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	var digimon models.DigimonDetail
	if err := json.NewDecoder(resp.Body).Decode(&digimon); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &digimon, nil
}
