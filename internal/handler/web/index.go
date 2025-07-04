package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/services"
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
	digimon, err := services.GetDigimonByName("agumon")
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
