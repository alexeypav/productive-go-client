package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"productive-go-client/internal/config"
	"productive-go-client/internal/data"
	"productive-go-client/internal/service"
	"productive-go-client/internal/ui"

	"github.com/cqroot/prompt"
)

var (
	serviceId = flag.String("serviceId", "", "ID of the service")
	date      = flag.String("date", time.Now().Format("2006-01-02"), "Date for the operation in YYYY-MM-DD format")
	hours     = flag.Int("hours", 8, "Hours component of the time")
	minutes   = flag.Int("minutes", 0, "Minutes component of the time")
	notes     = flag.String("notes", "", "Date for the operation in YYYY-MM-DD format")
	baseURL   = "https://api.productive.io/api/v2/"
)

func main() {
	fmt.Println("Checking Config")

	cfg, err := config.LoadAndVerifyConfig()
	if err != nil {
		log.Fatal("Error:", err)
	}

	client := &http.Client{}

	productiveService := data.NewProductive(client, baseURL)
	timeService := service.NewTimeService(productiveService)

	user, err := timeService.Repo.GetUser(cfg)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ui.DisplayUserDetails(&user)
	if err != nil {
		log.Fatalf("Error in %s \n", err)
	}

	if err := config.UpdateConfigWithUserID(&cfg, user.ID); err != nil {
		log.Print("Failed to update user ID:", err)
	}

	//If flags are passed handle and run, otherwise show the ui
	if len(os.Args) > 1 {
		flag.Parse()

		if *serviceId == "" {
			log.Fatal("No Service ID provided via flags...")
		}

		err = timeService.EnterTime(&user, &cfg, *date, *serviceId, *notes, *hours, *minutes)
		if err != nil {
			log.Fatal("Failed to create a time entry", err)
		}
		log.Printf("Successfully entered time for: Service ID: %v, Date: %v", *serviceId, *date)

	} else {

		for {
			choice, err := prompt.New().Ask("Choose:").Choose([]string{"Enter Time", "Show Time Codes", "Exit"})
			if err != nil {
				fmt.Println("Error:", err)
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
}
