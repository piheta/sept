package services

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
)

func SetUpKeys() error {
	if KeysExist() {
		return nil // don't make new keys
	}

	public_key, private_key, err := ed25519.GenerateKey(nil) // pass in nil so crypto/rand.Reader is used
	if err != nil {
		return fmt.Errorf("failed to create keypair")
	}

	err = saveKeyToFile(db.SEPT_DATA+"/private_key.pem", private_key, 0400)
	if err != nil {
		return fmt.Errorf("failed to save private key: %v", err)
	}

	err = saveKeyToFile(db.SEPT_DATA+"/public_key.pem", public_key, 0400)
	if err != nil {
		return fmt.Errorf("failed to save public key: %v", err)
	}

	return nil
}

func saveKeyToFile(filename string, key []byte, perm os.FileMode) error {
	err := os.WriteFile(filename, key, perm)
	if err != nil {
		return fmt.Errorf("failed to write key to file: %v", err)
	}

	// Set file permissions to 400 (read-only for the owner)
	err = os.Chmod(filename, perm)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %v", err)
	}

	return nil
}

func KeysExist() bool {
	_, err := os.Stat(db.SEPT_DATA + "/private_key.pem")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(db.SEPT_DATA + "/public_key.pem")
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func GetPublicKeyBase64() (string, error) {
	pubKey, err := os.ReadFile(db.SEPT_DATA + "/public_key.pem")
	if err != nil {
		return "", fmt.Errorf("failed to read public key: %v", err)
	}
	encoded := base64.StdEncoding.EncodeToString(pubKey)

	return encoded, nil
}

func SignMessage(message models.Message) (models.Message, error) {
	privateKeyFileContent, err := os.ReadFile(db.SEPT_DATA + "/private_key.pem")
	if err != nil {
		return models.Message{}, fmt.Errorf("failed to sign message: %v", err)
	}

	privateKey := ed25519.PrivateKey(privateKeyFileContent)

	dataToSign := fmt.Sprintf("%s|%s|%s", message.ChatID, message.UserID, message.Content)
	messageBytes := []byte(dataToSign)

	signature := ed25519.Sign(privateKey, messageBytes)

	message.Signature = base64.StdEncoding.EncodeToString(signature)

	return message, nil
}

func ExtractUserFromJwt(tokenString string) (models.User, error) {
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

func VerifyToken(tokenString string) error {
	publicKeyString, err := GetPublicKey()
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

func GetPublicKey() (string, error) {
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
