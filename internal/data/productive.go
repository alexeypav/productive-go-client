package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"productive-go-client/internal/models"
)

type Productive struct {
	Client *http.Client
	BaseURL string
}

func NewProductive(client *http.Client, baseURL string) *Productive {
	return &Productive{
		Client: client,
		BaseURL: baseURL,
	}
}


func (repo *Productive) GetUser(config models.Config) (models.User, error) {
    endpoint := repo.BaseURL + "people?filter[email]=" + fmt.Sprintf("%s", config.UserEmail)
    
    headers := map[string]string{
        "Content-Type":      "application/vnd.api+json",
        "X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
        "X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
    }

    // Create a new GET request
    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return models.User{}, err
    }

    // Add custom headers to the request
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    // Send the request using the repository's client
    resp, err := repo.Client.Do(req)
    if err != nil {
        return models.User{}, err
    }
    defer resp.Body.Close()

    // Decode JSON response
    var personResponse PersonResponse
    err = json.NewDecoder(resp.Body).Decode(&personResponse)
    if err != nil {
        return models.User{}, err
    }

    if len(personResponse.Data) == 0 {
        return models.User{}, fmt.Errorf("no person data found in the response")
    }

    return personResponse.Data[0], nil
}

// PersonResponse struct definition remains unchanged
type PersonResponse struct {
    Data []models.User `json:"data"`
}


func (repo *Productive) GetServiceAssignments(config models.Config) ([]models.ServiceAssignment, error) {
    endpoint := repo.BaseURL + "services"
    
    headers := map[string]string{
        "Content-Type":      "application/vnd.api+json",
        "X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
        "X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
    }

    // Create a new GET request
    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return nil, err
    }

    // Add custom headers to the request
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    // Send the request using the repository's client
    resp, err := repo.Client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Decode JSON response
    var serviceResponse ServiceResponse
    err = json.NewDecoder(resp.Body).Decode(&serviceResponse)
    if err != nil {
        return nil, err
    }

    if len(serviceResponse.Data) == 0 {
        return nil, fmt.Errorf("no service data found in the response")
    }

    return serviceResponse.Data, nil
}

// ServiceResponse struct remains unchanged
type ServiceResponse struct {
    Data []models.ServiceAssignment `json:"data"`
}


func (repo *Productive) PostTimeEntry(config models.Config, timeEntry models.TimeEntry) error {
    endpoint := repo.BaseURL + "/time_entries"
    
    // Create JSON request body
    headers := map[string]string{
        "Content-Type":      "application/vnd.api+json",
        "X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
        "X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
    }

    requestBody := map[string]interface{}{
        "data": timeEntry,
    }

    // Convert request body to JSON
    requestBodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        return err
    }

    // Send POST request with JSON body
    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        return err
    }

    // Set custom headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    // Send the HTTP request using the repository's client
    resp, err := repo.Client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check response status
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}
