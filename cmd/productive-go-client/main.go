package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"productive-go-client/internal/api"
	"productive-go-client/internal/models"
	"strconv"
	"strings"
)

var (
	baseURL = "https://api.productive.io/api/v2/"
)

func main() {

	fmt.Println("Checking Config")

	var config models.Config
	err := getConfig(&config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	endpoint := baseURL + "people?filter[email]=" + fmt.Sprintf("%s", config.UserEmail)
	headers := map[string]string{
		"Content-Type":      "application/vnd.api+json",
		"X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
		"X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
	}

	user, err := api.GetUser(endpoint, headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the user
	fmt.Println("User:")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("First Name: %s\n", user.Attributes.FirstName)
	fmt.Printf("Last Name: %s\n", user.Attributes.LastName)
	fmt.Printf("Email: %s\n", user.Attributes.Email)
	fmt.Printf("Title: %s\n", user.Attributes.Title)
	fmt.Println("----------------------------------")

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Get Time Codes")
		fmt.Println("2. Enter Time")
		fmt.Println("3. Option 3")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case 1:
			getAvailableTimeCodes(&config)
		case 2:
			enterTime(&user, &config)
		case 3:
			fmt.Println("You selected Option 3")
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func getAvailableTimeCodes(config *models.Config) {

	endpoint := baseURL + "services"
	headers := map[string]string{
		"Content-Type":      "application/vnd.api+json",
		"X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
		"X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
	}

	availableTimeCodes, err := api.GetServiceAssignments(endpoint, headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, assignment := range availableTimeCodes {
		fmt.Printf("ID: %s, Name: %s\n", assignment.ID, assignment.Attributes.FirstName)
	}
}

func enterTime(user *models.User, config *models.Config) {

	readInput("Enter Details") // FIX THIS IT SKIPS OVER FIRST PROMPT

	date := readInput("Enter Date, Format:")
	serviceID := readInput("Enter Service ID:")
	notes := readInput("Enter Notes:")
	time := readFloatInput("Enter Time in hours, part hours OK e.g 1.5:")

	//baseURL +
	endpoint := baseURL + "/time_entries"
	headers := map[string]string{
		"Content-Type":      "application/vnd.api+json",
		"X-Auth-Token":      fmt.Sprintf("%s", config.AccessToken),
		"X-Organization-Id": fmt.Sprintf("%s", config.CompanyId),
	}

	fmt.Printf("Endpoint: %s, Headers: %s, ServiceID: %s, Date: %s, UserID: %s, Notes: %s, Time: %f\n", endpoint, "headers", serviceID, date, user.ID, "", time)

	err := api.PostTimeEntry(endpoint, headers, serviceID, date, user.ID, notes, time)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

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
		err = saveConfig(*config)
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

func saveConfig(config models.Config) error {
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

func readInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func readFloatInput(prompt string) float64 {
	fmt.Println(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	time, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error converting input to float:", err)
		os.Exit(1)
	}
	return time
}
