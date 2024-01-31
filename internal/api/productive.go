package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"productive-go-client/internal/models"
)

var client = http.Client{}

func GetUser(endpoint string, headers map[string]string) (models.User, error) {
	// Create a new HTTP client

	// Create a new GET request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return models.User{}, err
	}

	// Add custom headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	resp, err := client.Do(req)
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

type PersonResponse struct {
	Data []models.User `json:"data"`
}

func GetServiceAssignments(endpoint string, headers map[string]string) ([]models.ServiceAssignment, error) {
	// Create a new HTTP client

	// Create a new GET request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add custom headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	resp, err := client.Do(req)
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
		return nil, fmt.Errorf("no person data found in the response")
	}

	return serviceResponse.Data, nil
}

type ServiceResponse struct {
	Data []models.ServiceAssignment `json:"data"`
}

func PostTimeEntry(url string, headers map[string]string, serviceId, dateInput, userId, notes string, time float64) error {
	// Create JSON request body

	requestBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "time_entries",
			"attributes": map[string]interface{}{
				"note": notes,
				"date": dateInput,
				"time": (time * 60),
			},
			"relationships": map[string]interface{}{
				"person": map[string]interface{}{
					"data": map[string]interface{}{
						"type": "people",
						"id":   userId,
					},
				},
				"service": map[string]interface{}{
					"data": map[string]interface{}{
						"type": "services",
						"id":   serviceId,
					},
				},
			},
		},
	}

	// Convert request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	// Convert byte slice to string
	requestBodyString := string(requestBodyBytes)

	// Print JSON request body
	fmt.Println("Request Body:", requestBodyString)

	// Send POST request with JSON body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}

	// Set content type header
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	// Print response body
	//fmt.Println("Response Body:", string(responseBody))

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
