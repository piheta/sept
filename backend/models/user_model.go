package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Ip        string `json:"ip"`
	Avatar    string `json:"avatar"`
	PublicKey string `josn:"public_key"`
}
