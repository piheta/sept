package models

type Message struct {
	ID        int    `json:"id"`
	ChatID    int    `json:"chat_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Signature string `json:"signature"`
}
