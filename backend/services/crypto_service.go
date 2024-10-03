package services

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"os"

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

	err = saveKeyToFile("./sept_data/private_key.pem", private_key, 0400)
	if err != nil {
		return fmt.Errorf("failed to save private key: %v", err)
	}

	err = saveKeyToFile("./sept_data/public_key.pem", public_key, 0400)
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
	_, err := os.Stat("./sept_data/private_key.pem")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat("./sept_data/public_key.pem")
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func GetPublicKeyBase64() (string, error) {
	pubKey, err := os.ReadFile("./sept_data/public_key.pem")
	if err != nil {
		return "", fmt.Errorf("failed to read public key: %v", err)
	}
	encoded := base64.StdEncoding.EncodeToString(pubKey)

	return encoded, nil
}

func SignMessage(message models.Message) (models.Message, error) {
	privateKeyFileContent, err := os.ReadFile("./sept_data/private_key.pem")
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
