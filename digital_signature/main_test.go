package main

import (
	"testing"
)

func TestDigitalSignature(t *testing.T) {

	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		t.Errorf("Error generating key pair: %s", err)
	}

	originalMessage := []byte("Hello, this is a test message.")

	signature, err := signMessage(originalMessage, privateKey)
	if err != nil {
		t.Errorf("Error signing the message: %s", err)
	}

	err = verifySignature(originalMessage, signature, publicKey)
	if err != nil {
		t.Errorf("Verification failed: %s", err)
	}

	alteredMessage := []byte("Altered message")
	err = verifySignature(alteredMessage, signature, publicKey)
	if err == nil {
		t.Errorf("Altered message verification succeeded, but it was expected to fail")
	}
}
