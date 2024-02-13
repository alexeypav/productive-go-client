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

	var productiveService data.ProductiveService
	productiveService = data.NewProductive(client, baseURL)
	timeService := service.NewTimeService(productiveService)

	user, err := timeService.Repo.GetUser(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ui.DisplayUserDetails(&user)
	if err != nil {
		fmt.Printf("Error in %s \n", err)
	}

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
			err := ui.EnterTime(&user, timeService, &cfg)
			if err != nil {
				fmt.Printf("Error in %s \n", err)
			}

		case "Show Time Codes":
			ui.ShowUserTimeCodes(timeService, &cfg)
			if err != nil {
				fmt.Printf("Error in %s \n", err)
			}

		case "Exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
