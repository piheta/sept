package services

import (
	"crypto/ed25519"
	"fmt"
)

func SetUpKeys() error {
	public_key, private_key, err := ed25519.GenerateKey(nil) // pass in nil so crypto/rand.Reader is used
	if err != nil {
		return fmt.Errorf("failed to create keypair")
	}

	fmt.Println(public_key)
	fmt.Println(private_key)

	return nil
}
