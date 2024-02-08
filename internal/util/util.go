package util

import (
	"encoding/json"
	"fmt"
	"productive-go-client/internal/models"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// For the prompt choose menu if chosing from a slice of structs
func StructListToStringList[T any](structList []T) ([]string, error) {
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

func getServiceAssignmentByID(serviceAssignments []models.ServiceAssignment, id string) models.ServiceAssignment {
	for _, sa := range serviceAssignments {
		if sa.Service_ID == id {
			return sa
		}
	}
	return models.ServiceAssignment{}
}
