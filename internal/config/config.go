package config

import (
	"encoding/json"
	"fmt"
	"os"

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

func getConfig(config *models.Config) error {
	// Check if the config file exists
	_, err := os.Stat("config.json")
	if os.IsNotExist(err) {
		// Config file does not exist, prompt for access token
		fmt.Print("Enter your access token: ")
		var accessToken string
		_, err := fmt.Scan(&accessToken)
		if err != nil {
			return err
		}

		// Save the access token to the config struct
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

		config.CompanyId = companyID

		// Save the config struct to the config file
		err = saveConfig(config)
		if err != nil {
			return err
		}
	} else {
		// Config file exists, load the config from it
		file, err := os.ReadFile("config.json")
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

	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
