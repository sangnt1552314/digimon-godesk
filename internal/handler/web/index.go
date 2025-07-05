package web

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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

func AddDigimonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	digimonName := strings.ToLower(r.FormValue("name"))
	if digimonName == "" {
		http.Error(w, "Digimon name is required", http.StatusBadRequest)
		return
	}

	digimon, err := services.GetDigimonByName(digimonName)
	if err != nil {
		http.Error(w, "Failed to fetch digimon", http.StatusInternalServerError)
		return
	}

	if digimon == nil {
		http.Error(w, "Digimon not found", http.StatusNotFound)
		return
	}

	tmplPath := "templates/index/fragments/digimon_response.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, digimon); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetDigimonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/digimon/")
	if id == "" {
		http.Error(w, "Digimon ID is required", http.StatusBadRequest)
		return
	}

	digimon, err := services.GetDigimonByName(id)
	if err != nil {
		http.Error(w, "Failed to fetch digimon", http.StatusInternalServerError)
		return
	}
	if digimon == nil {
		http.Error(w, "Digimon not found", http.StatusNotFound)
		return
	}

	tmplPath := "templates/digimon/detail.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, digimon); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
