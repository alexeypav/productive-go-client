package service

import (
	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

func LoadAndVerifyConfig() (models.Config, error) {
	return data.LoadConfig()
}

func UpdateOrCreateConfig(accessToken, companyID, userEmail string) error {
	config := models.Config{
		AccessToken: accessToken,
		CompanyId:   companyID,
		UserEmail:   userEmail,
	}

	return data.SaveConfig(&config)
}
