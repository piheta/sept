package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
)

var AuthedUser = models.User{}

type AuthService struct {
}

func NewAuthSerivce() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Login(email, password string) (*models.User, error) {
	// Generate keys, if they dont exist
	if err := SetUpKeys(); err != nil {
		return nil, fmt.Errorf("error setting up keys: %w", err)
	}

	public_key, err := GetPublicKeyBase64()
	if err != nil {
		return nil, fmt.Errorf("failed to get the client's public key: %w", err)
	}

	data := map[string]string{
		"email":      email,
		"password":   password,
		"public_key": public_key,
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

	user, err := ExtractUserFromJwt(token)
	if err != nil {
		return nil, fmt.Errorf("token not found in response")
	}

	if err := VerifyToken(token); err != nil {
		return nil, fmt.Errorf("failed to verify token")
	}

	if err := as.saveJwt(token); err != nil {
		return nil, fmt.Errorf("error saving jwt to file")
	}

	user.PublicKey, err = GetPublicKeyBase64()
	if err != nil {
		log.Fatalf("failed to get public key")
	}
	db.InitDb(user)

	AuthedUser = user
	go SnConnectionHandler()

	return &user, nil
}

func (as *AuthService) Register(username, email, password string) (*map[string]interface{}, error) {
	fmt.Println(username, email)
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

func (as *AuthService) saveJwt(content string) error {
	file, err := os.Create(db.SEPT_DATA + "/user.jwt")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Content to write to the file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func (as *AuthService) LogInWithExistingJwt() error {
	jwt, err := os.ReadFile(db.SEPT_DATA + "/user.jwt")
	if err != nil {
		return fmt.Errorf("jwt does not exist %w ", err)
	}
	jwtString := string(jwt)

	if err := VerifyToken(jwtString); err != nil {
		as.LogOut()
		return fmt.Errorf("jwt is not valid %w, ", err)
	}

	if !KeysExist() {
		return fmt.Errorf("failed to log in with jwt %w, ", err)
	}

	user, _ := ExtractUserFromJwt(jwtString)
	user.PublicKey, err = GetPublicKeyBase64()
	if err != nil {
		log.Fatalf("failed to get public key")
	}

	if err := db.InitDb(user); err != nil {
		return fmt.Errorf("failed to init db with jwt %w ", err)
	}

	AuthedUser = user
	go SnConnectionHandler()

	return nil
}

func (as *AuthService) LogOut() error {
	err := os.Remove(db.SEPT_DATA + "/user.jwt")
	if err != nil {
		return fmt.Errorf("failed to log out, can't delete jwt %w ", err)
	}

	return nil
}
