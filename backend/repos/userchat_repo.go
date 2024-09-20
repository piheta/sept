package repos

import (
	"database/sql"
)

type UserchatRepo struct {
	db *sql.DB
}

func NewUserchatRepo(db *sql.DB) *UserchatRepo {
	return &UserchatRepo{db: db}
}

func (ucr *UserchatRepo) AddUserToChat(user_id string, chat_id string) error {
	query := `INSERT OR IGNORE INTO user_chats (user_id, chat_id) VALUES (?, ?)`
	_, err := ucr.db.Exec(query, user_id, chat_id)
	return err
}
