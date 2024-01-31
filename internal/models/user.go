package models

//Map to response from api - keeping Attributes node for simplicity
type User struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RoleID    int    `json:"role_id"`
		Email     string `json:"email"`
		Title     string `json:"title"`
		UserID    int    `json:"user_id"`
	} `json:"attributes"`
}
