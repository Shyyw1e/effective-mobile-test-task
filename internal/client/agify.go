package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"
)

type AgifyResponse struct {
	Age int `json:"age"`
}

func GetAge(name string) (int, error) {
	logger.Log.Debug("Calling Agify API", "name", name)

	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Error("Failed to request Agify", "err", err)
		return 0, err
	}
	defer resp.Body.Close()

	var result AgifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Log.Error("Failed to decode Agify response", "err", err)
		return 0, err
	}

	logger.Log.Debug("Agify API response", "age", result.Age)
	return result.Age, nil
}
