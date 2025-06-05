package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"
)

type GenderizeResonse struct {
	Gender string	`json:"gender"`
}

func GetGender(name string) (string, error) {
	logger.Log.Debug("Calling Genderize API", "name", name)

	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Error("Failed to request Genderize", "err", err)
		return "", err
	}
	defer resp.Body.Close()

	var result GenderizeResonse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Log.Error("Failed to decode Genderize response", "err", err)
		return "", err
	}

	logger.Log.Debug("Genderize API response", "gender", result.Gender)
	return result.Gender, nil
}