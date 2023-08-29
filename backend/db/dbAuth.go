package db

import (
	"database/sql"
	"log"
)

func GetSaltOfUser(email string) (string, error) {
	rawSQL := `SELECT salt FROM users WHERE email = $1`
	row := db.QueryRow(rawSQL, email)
	var salt string
	err := row.Scan(&salt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("User not found")
		} else {
			log.Print(err)
			return "", err
		}
	}
	return salt, nil
}

func Auth(sessionKey string) (bool, error) {
	log.Println("Enter db.Auth")
	defer log.Println("Exited db.Auth")
	rawSQL := `
		SELECT COUNT(*) AS session_count
		FROM user_session
		WHERE session_key = $1;
	`
	var count int

	err := db.QueryRow(rawSQL, sessionKey).Scan(&count)
	if err != nil {
		log.Print(err)
		return false, err
	}

	if count == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
