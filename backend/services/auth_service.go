package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
)

type AuthService struct {
	user_repo     *repos.UserRepo
	chat_repo     *repos.ChatRepo
	userchat_repo *repos.UserchatRepo
}

func NewAuthSerivce(userRepo *repos.UserRepo, chatRepo *repos.ChatRepo, userchatRepo *repos.UserchatRepo) *AuthService {
	return &AuthService{
		user_repo:     userRepo,
		chat_repo:     chatRepo,
		userchat_repo: userchatRepo,
	}
}

func (as *AuthService) Login(email, password string) (*models.User, error) {
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

	user, err := as.extractUserFromUnverifiedClaims(token)
	if err != nil {
		return nil, fmt.Errorf("token not found in response")
	}

	if err := as.VerifyToken(token); err != nil {
		return nil, fmt.Errorf("failed to verify token")
	}

	//FIRST TIME LOGIN
	db.InitDb(user.ID) // creates db and salt file for future encryption.
	//CREATE KEYPAIR
	// SetUpKeys()
	as.user_repo.AddUser(user)
	as.chat_repo.AddChat(user.Username) // Create chat named the same as the username of the user
	chat, _ := as.chat_repo.GetChatByName(user.Username)
	as.userchat_repo.AddUserToChat(user.ID, chat.ID) // link user with the chat

	return &user, nil
}

func (as *AuthService) Register(username, email, password string) (*map[string]interface{}, error) {
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

func (as *AuthService) extractUserFromUnverifiedClaims(tokenString string) (models.User, error) {
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
		ID:        user_id,
		Username:  username,
		Ip:        "127.0.0.1",
		Avatar:    "https://fuibax.github.io/images/fulls/knight_sylvia.png",
		PublicKey: "",
	}

	return user, nil
}

func (as *AuthService) VerifyToken(tokenString string) error {
	publicKeyString, err := as.GetPublicKey()
	if err != nil {
		return fmt.Errorf("error retrieving the public key: %w", err)
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM([]byte(publicKeyString))
	if err != nil {
		return fmt.Errorf("error parsing public key: %w", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok { // Verify the signing method
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is not valid")
	}

	return nil // Token is valid
}

func (as *AuthService) GetPublicKey() (string, error) {
	resp, err := http.Get("http://localhost:8080/api/key")
	if err != nil {
		return "", fmt.Errorf("error fetching public key: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var keyResp models.PublicKeyResponse
	err = json.Unmarshal(body, &keyResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return keyResp.PublicKey, nil
}
