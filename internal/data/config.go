// data/config_data.go
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"productive-go-client/internal/models"
	"strings"
)

func GetConfigFilePath() (filePath string, err error) {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return "", err
	}

	execDir := filepath.Dir(execPath)
	//If using go run, use invocation dir for config load/save
	if strings.Contains(execPath, "/tmp/") || strings.Contains(execPath, `\Temp\`) || strings.Contains(execPath, `/T/`) {
		fmt.Println("Program detected to be likely run using 'go run', using invocation dir for config file.")
		filePath = "config.json"
	} else {
		filePath = filepath.Join(execDir, "config.json")
	}
	return filePath, nil
}

func LoadConfig() (models.Config, error) {
	var config models.Config

	filePath, err := GetConfigFilePath()
	if err != nil {
		return config, err
	}

	file, err := os.ReadFile(filePath)
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

	filePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, file, 0644)
}
