package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"
)

type Nationality struct {
	CountryID		string			`json:"country_id"`
	Probability		float64			`json:"probability"`
}

type NationalizeResponse struct {
	Country 		[]Nationality	`json:"country"`	
}

func GetNationalities(name string) ([]string, error) {
	logger.Log.Debug("Calling Nationalize API", "name", name)

	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Error("Failed to Request Nationalize", "err", err)
		return []string{}, err
	}
	defer resp.Body.Close()

	var result NationalizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Log.Error("Failed to decode response", "err", err)
		return []string{}, err
	}

	nationalities := make([]string, 0, len(result.Country))
	for _, country := range result.Country {
		nationalities = append(nationalities, country.CountryID)
	}

	return nationalities, nil
}