package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"productive-go-client/internal/api"
	"productive-go-client/internal/models"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func main() {

	fmt.Println("Checking Config")

	var config models.Config
	err := getConfig(&config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user, err := api.GetUser(config)
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

	config.UserId = user.ID
	err = saveConfig(&config)
	if err != nil {
		fmt.Println("Failed to update user ID:", err)
		return
	}

	for {
		//Display Menu loop
		choice, err := prompt.New().Ask("Choose:").
			Choose([]string{"Enter Time", "Show Time Codes", "Exit"})
		checkErr(err)

		switch choice {
		case "Enter Time":
			enterTime(&user, &config)
		case "Show Time Codes":
			availableTimeCodes, err := api.GetServiceAssignments(config)
			checkErr(err)
			printList(availableTimeCodes)

		case "Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func enterTime(user *models.User, config *models.Config) {

	//Date
	today := time.Now().Format("2006-01-02")
	date, err := prompt.New().Ask("Enter date:").Input(today)
	checkErr(err)

	//Time code
	availableTimeCodes, err := api.GetServiceAssignments(*config)
	if err != nil {
		fmt.Printf("Unable to save the customers to a file: %s", err.Error())
		return
	}
	availableTimeCodesString, err := structListToStringList(availableTimeCodes)
	if err != nil {
		fmt.Printf("Unable to save the customers to a file: %s", err.Error())
		return
	}
	printList(availableTimeCodes)
	serviceAssignment, err := prompt.New().Ask("Choose:").
		Choose(availableTimeCodesString)
	checkErr(err)
	fmt.Print(serviceAssignment)
	serviceID, err := prompt.New().Ask("Enter Time Code ID:").
		Input("", input.WithInputMode(input.InputInteger))
	checkErr(err)

	//Notes
	notes, err := prompt.New().Ask("Enter any notes:").Input("")
	checkErr(err)

	//Hours
	timeH, err := prompt.New().Ask("Enter Time (Hours):").Input("0", input.WithInputMode(input.InputInteger))
	checkErr(err)
	timeHours, err := strconv.Atoi(timeH)
	checkErr(err)

	//Minutes
	timeM, err := prompt.New().Ask("Enter Time (Minutes):").Input("0", input.WithInputMode(input.InputInteger))
	checkErr(err)
	timeMinutes, err := strconv.Atoi(timeM)
	checkErr(err)

	timeMinutes = timeMinutes + timeHours*60

	//Selection result
	fmt.Printf(" ServiceID: %v, Date: %s, UserID: %s, Notes: %s, Time: %v\n", serviceID, date, user.ID, notes, timeMinutes)

	_, err = fmt.Println("Done") //api.PostTimeEntry(*config, serviceAssignment.Service_ID, date, user.ID, notes, timeMinutes)
	if err != nil {
		fmt.Println("Error posting time entry:", err)
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

func getServiceAssignmentByID(serviceAssignments []models.ServiceAssignment, id string) models.ServiceAssignment {
	for _, sa := range serviceAssignments {
		if sa.Service_ID == id {
			return sa
		}
	}
	return models.ServiceAssignment{}
}

func checkErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func printList[T any](list []T) {
	fmt.Println("----------------------------------")
	for _, item := range list {
		fmt.Println(item)
	}
	fmt.Println("----------------------------------")
}

// For the prompt choose menu if chosing from a slice of structs
func structListToStringList[T any](structList []T) ([]string, error) {
	var stringList []string
	for _, sl := range structList {
		jsonBytes, err := json.Marshal(sl)
		if err != nil {
			return stringList, err
		}
		stringList = append(stringList, string(jsonBytes))
	}
	return stringList, nil
}
