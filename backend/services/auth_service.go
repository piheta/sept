package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
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

	user, err := as.ExtractUserFromJwt(token)
	if err != nil {
		return nil, fmt.Errorf("token not found in response")
	}

	if err := as.VerifyToken(token); err != nil {
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

func (as *AuthService) ExtractUserFromJwt(tokenString string) (models.User, error) {
	var user_id string
	var username string
	var user_ip string
	var public_key string
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

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user_ip = fmt.Sprint(claims["ip"])
	}
	if user_ip == "" {
		return models.User{}, fmt.Errorf("invalid token payload")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		public_key = fmt.Sprint(claims["public_key"])
	}
	if public_key == "" {
		return models.User{}, fmt.Errorf("invalid token payload")
	}

	user := models.User{
		ID:        user_id,
		Username:  username,
		Ip:        user_ip,
		Avatar:    "https://fuibax.github.io/images/fulls/knight_sylvia.png",
		PublicKey: public_key,
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

func (as *AuthService) saveJwt(content string) error {
	file, err := os.Create("./sept_data/user.jwt")
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
	jwt, err := os.ReadFile("./sept_data/user.jwt")
	if err != nil {
		return fmt.Errorf("jwt does not exist %w ", err)
	}
	jwtString := string(jwt)

	if err := as.VerifyToken(jwtString); err != nil {
		as.LogOut()
		return fmt.Errorf("jwt is not valid %w, ", err)
	}

	if !KeysExist() {
		return fmt.Errorf("failed to log in with jwt %w, ", err)
	}

	user, _ := as.ExtractUserFromJwt(jwtString)
	user.PublicKey, err = GetPublicKeyBase64()
	if err != nil {
		log.Fatalf("failed to get public key")
	}

	if err := db.InitDb(user); err != nil {
		return fmt.Errorf("failed to init db with jwt %w ", err)
	}

	AuthedUser = user

	return nil
}

func (as *AuthService) LogOut() error {
	err := os.Remove("./sept_data/user.jwt")
	if err != nil {
		return fmt.Errorf("failed to log out, can't delete jwt %w ", err)
	}

	return nil
}
