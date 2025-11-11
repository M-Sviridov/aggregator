package config

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}
