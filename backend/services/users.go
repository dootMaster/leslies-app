package services

import (
	"crypto/rand"
	"encoding/hex"
	"leslies-app/backend/db"
	"leslies-app/backend/shared"
	"log"
)

func CreateUser(user shared.CreateUserArgs) error {
	log.Println("Entered services.CreateUser")
	defer log.Println("Exited services.CreateUser")

	salt, err := GenerateSalt()
	if err != nil {
		log.Print(err)
		return err
	}

	user.Password = HashPassword(user.Password + salt)

	err = db.CreateUser(user, salt)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func generateSessionKey() (string, error) {
	sessionKey := make([]byte, 72)
	_, err := rand.Read(sessionKey)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return hex.EncodeToString(sessionKey), nil
}

func Login(credentials shared.CredentialArgs) (string, error) {
	log.Println("Entered services.fmtin")
	defer log.Println("Exited services.fmtin")
	sessionKey, err := generateSessionKey()
	if err != nil {
		log.Print(err)
		return "", err
	}

	salt, err := db.GetSaltOfUser(credentials.Email)
	if err != nil {
		log.Print(err)
		return "", err
	}
	credentials.Password = HashPassword(credentials.Password + salt)

	err = db.Login(credentials, sessionKey)

	if err != nil {
		log.Print(err)
		return "", err
	}
	return sessionKey, nil
}

func Logout(sessionKey string, logoutAllSessions bool) error {
	log.Println("Entered services.fmtin")
	defer log.Println("Exited services.fmtin")

	err := db.Logout(sessionKey, logoutAllSessions)

	return err
}
