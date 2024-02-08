package ui

import (
	"fmt"

	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

func FetchAndDisplayUser(config models.Config) (models.User, error) {
	user, err := data.GetUser(config)
	if err != nil {
		return models.User{}, err
	}
	displayUserDetails(user)
	return user, nil
}

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
