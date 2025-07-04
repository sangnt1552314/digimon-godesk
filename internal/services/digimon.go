package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sangnt1552314/digimon-godesk/internal/models"
)

const (
	baseURL = "http://digi-api.com/api/v1/digimon"
)

func GetDigimonByName(name string) (*models.DigimonDetail, error) {
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
