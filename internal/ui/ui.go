package ui

import (
	"encoding/json"
	"fmt"
	"productive-go-client/internal/models"
	"productive-go-client/internal/service"
	"productive-go-client/internal/util"
	"strconv"
	"time"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func EnterTime(user *models.User, timeService *service.TimeService, config *models.Config) error {
	// Get available time codes
	availableTimeCodes, err := timeService.GetServiceAssignments(*config)
	if err != nil {
		fmt.Printf("Error fetching available time codes: %s\n", err.Error())
		return err
	}

	// Convert to []string for the choose prompt
	availableTimeCodesString, err := util.StructListToStringList(availableTimeCodes)
	if err != nil {
		fmt.Printf("Error converting time codes to string list: %s\n", err.Error())
		return err
	}

	// Prompt user for inputs...
	// For simplicity, we'll assume you have functions similar to `prompt.New().Ask(...)` for user interaction.
	today := time.Now().Format("2006-01-02")
	date, err := prompt.New().Ask("Enter Date:").Input(today)

	serviceAssignmentString, err := prompt.New().Ask("Choose Time Code:").
		Choose(availableTimeCodesString)
	var serviceAssignment models.ServiceAssignment
	err = json.Unmarshal([]byte(serviceAssignmentString), &serviceAssignment)
	serviceAssignmentID := serviceAssignment.Service_ID

	notes, err := prompt.New().Ask("Enter Notes for Time Entry:").Input("")

	//Hours
	timeH, err := prompt.New().Ask("Enter Time (Hours):").Input("0", input.WithInputMode(input.InputInteger))
	util.CheckErr(err)
	timeHours, err := strconv.Atoi(timeH)
	util.CheckErr(err)

	//Minutes
	timeM, err := prompt.New().Ask("Enter Time (Minutes):").Input("0", input.WithInputMode(input.InputInteger))
	util.CheckErr(err)
	timeMinutes, err := strconv.Atoi(timeM)
	util.CheckErr(err)

	// Call business logic layer to process and create time entry
	err = timeService.EnterTime(user, config, date, serviceAssignmentID, notes, timeHours, timeMinutes)
	if err != nil {
		fmt.Println("Error posting time entry:", err)
		return err
	}

	fmt.Println("Time Entered")
	return nil
}

// Implement promptForDate, promptForServiceAssignment, promptForNotes, and promptForTime as needed.



// func FetchAndDisplayUser(config models.Config) (models.User, error) {
// 	user, err := data.GetUser(config)
// 	if err != nil {
// 		return models.User{}, err
// 	}
// 	displayUserDetails(user)
// 	return user, nil
// }

func displayUserDetails(user models.User) {
	// Print the user details
	fmt.Println("User:")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("First Name: %s\n", user.Attributes.FirstName)
	fmt.Printf("Last Name: %s\n", user.Attributes.LastName)
	fmt.Printf("Email: %s\n", user.Attributes.Email)
	fmt.Printf("Title: %s\n", user.Attributes.Title)
	fmt.Println("----------------------------------")
}

// Print struct to screen for display, user, config etc.
func PrintList[T any](list []T) {
	fmt.Println("----------------------------------")
	for _, item := range list {
		fmt.Println(item)
	}
	fmt.Println("----------------------------------")
}
