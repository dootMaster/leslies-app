package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
