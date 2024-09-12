package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/piheta/sept/infra/db"
	"github.com/piheta/sept/models"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	//get latest version
}

func (a *App) GetUsers() ([]models.User, error) {
	users, err := db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return users, nil
}

func (a *App) SendMessage(message string, chat_id int) ([]models.Message, error) {
	err := db.AddMessage(chat_id, 1, message)
	if err != nil {
		return nil, fmt.Errorf("failed to add message: %w", err)
	}
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetChatMessages(chat_id int) ([]models.Message, error) {
	return db.GetMessagesByChatID(chat_id)
}

func (a *App) GetUser(user_id int) (models.User, error) {
	return db.GetUser(user_id)
}

func (a *App) Login(email, password string) (*models.User, error) {
	data := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("registration failed, status code: %d", resp.StatusCode)
	}

	var res map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	token, ok := res["token"].(string)
	if !ok {
		return nil, fmt.Errorf("token not found in response")
	}

	user, err := extractUserFromUnverifiedClaims(token)
	if err != nil {
		return nil, fmt.Errorf("token not found in response")
	}

	db.InitDb(user.UserID, password) // creates db and salt file for future encryption.
	db.AddUser(user)
	db.AddChat(user.Username) // Create chat named the same as the username of the user
	chat, _ := db.GetChatByName(user.Username)
	db.AddUserToChat(user.UserID, chat.ID) // link user with the chat

	return &user, nil
}

func (a *App) Register(username, email, password string) (*map[string]interface{}, error) {
	fmt.Println(username, email, password)
	data := map[string]string{
		"name":     username,
		"email":    email,
		"password": password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/api/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("registration failed, status code: %d", resp.StatusCode)
	}

	var res map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &res, nil
}

func extractUserFromUnverifiedClaims(tokenString string) (models.User, error) {
	var user_id string
	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return models.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["name"])
	}
	if username == "" {
		return models.User{}, fmt.Errorf("invalid token payload")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user_id = fmt.Sprint(claims["id"])
	}
	if user_id == "" {
		return models.User{}, fmt.Errorf("invalid token payload")
	}

	user := models.User{
		UserID:   user_id,
		Username: username,
		Ip:       "127.0.0.1",
		Avatar:   "https://fuibax.github.io/images/fulls/knight_sylvia.png",
	}

	return user, nil
}
