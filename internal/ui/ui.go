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
	// Convert to []string for the ui choose prompt
	availableTimeCodesString, err := util.StructListToStringList(availableTimeCodes)
	if err != nil {
		fmt.Printf("Error converting time codes to string list: %s\n", err.Error())
		return err
	}

	//Start prompts

	//Date
	today := time.Now().Format("2006-01-02")
	date, err := prompt.New().Ask("Enter Date:").Input(today)
	if err != nil {
		return fmt.Errorf("date: %w", err)
	}

	//Time Code
	serviceAssignmentString, err := prompt.New().Ask("Choose Time Code:").
		Choose(availableTimeCodesString)
	if err != nil {
		return fmt.Errorf("time code: %w", err)
	}

	//Selected time code back models.ServiceAssignment
	var serviceAssignment models.ServiceAssignment
	err = json.Unmarshal([]byte(serviceAssignmentString), &serviceAssignment)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}
	serviceAssignmentID := serviceAssignment.Service_ID

	// Notes
	notes, err := prompt.New().Ask("Enter Notes for Time Entry:").Input("")
	if err != nil {
		return fmt.Errorf("notes: %w", err)
	}
	//Hours
	timeH, err := prompt.New().Ask("Enter Time (Hours):").Input("0", input.WithInputMode(input.InputInteger))
	if err != nil {
		return fmt.Errorf("time hours: %w", err)
	}
	timeHours, err := strconv.Atoi(timeH)

	//Minutes
	timeM, err := prompt.New().Ask("Enter Time (Minutes):").Input("0", input.WithInputMode(input.InputInteger))
	if err != nil {
		return fmt.Errorf("time minutes: %w", err)
	}
	timeMinutes, err := strconv.Atoi(timeM)
	if err != nil {
		return fmt.Errorf("time minutes: %w", err)
	}

	//Create time entry
	err = timeService.EnterTime(user, config, date, serviceAssignmentID, notes, timeHours, timeMinutes)
	if err != nil {
		fmt.Println("Error posting time entry:", err)
		return err
	}

	fmt.Println("Time Entered")
	return nil
}

func ShowUserTimeCodes(timeService *service.TimeService, config *models.Config) error {

	availableTimeCodes, err := timeService.GetServiceAssignments(*config)
	if err != nil {
		fmt.Printf("Error fetching available time codes: %s\n", err.Error())
		return err
	}

	printStruct(availableTimeCodes)

	return nil

}

func DisplayUserDetails(user *models.User) {
	// Print the formatted user details
	fmt.Println("----------------------------------")
	fmt.Println("User Details")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("First Name: %s\n", user.Attributes.FirstName)
	fmt.Printf("Last Name: %s\n", user.Attributes.LastName)
	fmt.Printf("Email: %s\n", user.Attributes.Email)
	fmt.Printf("Title: %s\n", user.Attributes.Title)
	fmt.Println("----------------------------------")
}

// Print struct to screen for nicer display: some response, config etc.
func printStruct[T any](list []T) {
	fmt.Println("----------------------------------")
	for _, item := range list {
		fmt.Println(item)
	}
	fmt.Println("----------------------------------")
}
