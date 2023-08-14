package common

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken(length int) (string, error) {
	tokenBytes := make([]byte, length)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(tokenBytes), nil
}
