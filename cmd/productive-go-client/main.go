package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"productive-go-client/internal/config"
	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
	"productive-go-client/internal/ui"
	"productive-go-client/internal/util"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func main() {
	fmt.Println("Checking Config")

	cfg, err := config.LoadAndVerifyConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user, err := ui.FetchAndDisplayUser(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := config.UpdateConfigWithUserID(&cfg, user.ID); err != nil {
		fmt.Println("Failed to update user ID:", err)
		return
	}

	runMainMenu(&user, &cfg)
}

func runMainMenu(user *models.User, config *models.Config) {
	for {
		choice, err := prompt.New().Ask("Choose:").Choose([]string{"Enter Time", "Show Time Codes", "Exit"})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch choice {
		case "Enter Time":
			enterTime(user, config)
		case "Show Time Codes":
			handleShowTimeCodes(config)
		case "Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func handleShowTimeCodes(config *models.Config) {
	availableTimeCodes, err := data.GetServiceAssignments(*config)
	if err != nil {
		fmt.Println("Error retrieving time codes:", err)
		return
	}
	ui.PrintList(availableTimeCodes)
}

func enterTime(user *models.User, config *models.Config) error {

	//Get available time codes, convert to []string for the choose prompt
	availableTimeCodes, err := data.GetServiceAssignments(*config)
	if err != nil {
		fmt.Printf("Unable to save the customers to a file: %s", err.Error())
		return err
	}
	availableTimeCodesString, err := util.StructListToStringList(availableTimeCodes)
	if err != nil {
		fmt.Printf("Unable to save the customers to a file: %s", err.Error())
		return err
	}

	timeEntry := models.NewTimeEntry()
	fmt.Printf("%+v\n", timeEntry)

	//Date
	today := time.Now().Format("2006-01-02")
	timeEntry.Attributes.Date, err = prompt.New().Ask("Enter Date:").Input(today)
	util.CheckErr(err)

	//Time Code
	serviceAssignmentString, err := prompt.New().Ask("Choose Time Code:").
		Choose(availableTimeCodesString)
	util.CheckErr(err)
	//Back to struct to extract and use the ID
	var serviceAssignment models.ServiceAssignment
	err = json.Unmarshal([]byte(serviceAssignmentString), &serviceAssignment)
	timeEntry.Relationships.Service.Data.ID = serviceAssignment.Service_ID

	//Notes
	notes, err := prompt.New().Ask("Enter Notes for Time Entry:").Input("")
	util.CheckErr(err)
	timeEntry.Attributes.Note = notes

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

	timeEntry.Attributes.Time = timeMinutes + timeHours*60

	//Set User
	timeEntry.Relationships.Person.Data.ID = user.ID

	//Selection result
	fmt.Printf("%v \n", timeEntry)
	err = data.PostTimeEntry(*config, timeEntry)
	if err != nil {
		fmt.Println("Error posting time entry:", err)
		return err
	}

	fmt.Println("Time Entered") //

	return nil

}
