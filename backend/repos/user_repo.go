package repos

import (
	"database/sql"

	"github.com/piheta/sept/backend/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) GetUser(id string) (models.User, error) {
	query := `
        SELECT id, username, ip, avatar, public_key
        FROM users
        WHERE id = ?
    `
	var user models.User
	err := ur.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Ip, &user.Avatar, &user.PublicKey)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepo) GetUsers() ([]models.User, error) {
	query := `
		SELECT id, username, ip, avatar, public_key
		FROM users
	`
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Ip, &user.Avatar, &user.PublicKey); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) AddUser(user models.User) error {
	query := `INSERT OR IGNORE INTO users (id, username, ip, avatar, public_key) VALUES (?, ?, ?, ?, ?)`
	_, err := ur.db.Exec(query, user.ID, user.Username, user.Ip, user.Avatar, user.PublicKey)
	return err
}

func (ur *UserRepo) SetDB(db *sql.DB) {
	ur.db = db
}
