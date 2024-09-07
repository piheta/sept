package models

type User_model struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Ip       string `json:"ip"`
	Avatar   string `json:"avatar"`
}
