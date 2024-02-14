// data/config_data.go
package data

import (
	"encoding/json"
	"os"
	"productive-go-client/internal/models"
)

func LoadConfig() (models.Config, error) {
	var config models.Config

	file, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	return config, err
}

func SaveConfig(config *models.Config) error {
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", file, 0644)
}
