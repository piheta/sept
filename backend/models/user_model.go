package models

type User struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Ip        string `json:"ip"`
	Avatar    string `json:"avatar"`
	PublicKey string `josn:"public_key"`
}
