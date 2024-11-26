package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"go-leveldb/models"
)

func LoadLaunchesFromJSON(filepath string) ([]models.Launch, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	var launches []models.Launch
	if err := json.Unmarshal(data, &launches); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return launches, nil
}
