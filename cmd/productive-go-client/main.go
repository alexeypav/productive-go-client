package main

import (
	"fmt"
	"net/http"

	"productive-go-client/internal/config"
	"productive-go-client/internal/data"
	"productive-go-client/internal/service"
	"productive-go-client/internal/ui"

	"github.com/cqroot/prompt"
)

func main() {
	fmt.Println("Checking Config")

	cfg, err := config.LoadAndVerifyConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	baseURL := "https://api.productive.io/api/v2/"

	timeEntryRepo := data.NewProductive(client, baseURL)
	timeService := service.NewTimeService(timeEntryRepo)

	user, err := timeService.Repo.GetUser(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%v \n", user)

	if err := config.UpdateConfigWithUserID(&cfg, user.ID); err != nil {
		fmt.Println("Failed to update user ID:", err)
		return
	}

	for {
		choice, err := prompt.New().Ask("Choose:").Choose([]string{"Enter Time", "Show Time Codes", "Exit"})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch choice {
		case "Enter Time":
			ui.EnterTime(&user, timeService, &cfg)
		// case "Show Time Codes":
		// 	handleShowTimeCodes(config)
		case "Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// func handleShowTimeCodes(config *models.Config) {
// 	availableTimeCodes, err := data.GetServiceAssignments(*config)
// 	if err != nil {
// 		fmt.Println("Error retrieving time codes:", err)
// 		return
// 	}
// 	ui.PrintList(availableTimeCodes)
// }
