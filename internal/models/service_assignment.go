package models

//Map to response from api - keeping Attributes node for simplicity
type ServiceAssignment struct {
	Service_ID string `json:"id"`
	Attributes struct {
		FirstName string `json:"name"`
	} `json:"attributes"`
}
