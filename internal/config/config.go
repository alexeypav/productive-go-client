package config

import (
	"encoding/json"
	"fmt"
	"os"

	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

func LoadAndVerifyConfig() (models.Config, error) {
	var config models.Config
	err := getConfig(&config)
	return config, err
}

func UpdateConfigWithUserID(config *models.Config, userID string) error {
	config.UserId = userID
	return saveConfig(config)
}

// In Progress, Split and move promts to UI layer
func getConfig(config *models.Config) error {
	// Check if the config file exists
	configPath, err := data.GetConfigFilePath()
	if err != nil {
		return err
	}
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {

		fmt.Print("Enter your access token: ")
		var accessToken string
		_, err := fmt.Scan(&accessToken)
		if err != nil {
			return err
		}

		config.AccessToken = accessToken

		fmt.Print("Enter your company ID: ")
		var companyID string
		_, err = fmt.Scan(&companyID)
		if err != nil {
			return err
		}

		config.CompanyId = companyID

		fmt.Print("Enter your Email: ")
		var userEmail string
		_, err = fmt.Scan(&userEmail)
		if err != nil {
			return err
		}

		config.UserEmail = userEmail

		// Save the config to file
		err = saveConfig(config)
		if err != nil {
			return err
		}
	} else {
		// Config file exists, load the config from it
		file, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(file, config)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveConfig(config *models.Config) error {
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	configPath, err := data.GetConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
