package models

type Message struct {
	ID        int    `json:"id"`
	ChatID    string `json:"chat_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Signature string `json:"signature"`
}
