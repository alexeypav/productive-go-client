package models

type Config struct {
	AccessToken string `json:"access_token"`
	CompanyId   string `json:"company_id"`
	UserId      string `json:"user_id"`
	UserEmail   string `json:"user_email"`
}
