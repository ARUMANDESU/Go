package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func signMessage(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {

	hash := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func verifySignature(message, signature []byte, publicKey *rsa.PublicKey) error {
	hash := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return
	}

	originalMessage := []byte("Hello, this is a test message.")

	signature, err := signMessage(originalMessage, privateKey)
	if err != nil {
		fmt.Println("Error signing the message:", err)
		return
	}

	err = verifySignature(originalMessage, signature, publicKey)
	if err != nil {
		fmt.Println("Verification failed:", err)
		return
	}

	fmt.Println("Signature verified. Message is authentic.")
}
