package db

import (
	"fmt"
	"leslies-app/backend/shared"
	"log"
)

func CreateUser(user shared.CreateUserArgs, salt string) error {
	log.Print("Enter db.CreateUser")
	defer log.Print("Exited db.CreateUser")
	rawSQL := `
		INSERT INTO users (first_name, last_name, email, pw, salt)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (email) DO NOTHING;
	`

	result, err := db.Exec(
		rawSQL,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		salt,
	)
	if err != nil {
		log.Print(err)
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if affectedRows > 0 {
		log.Printf("user %s created", user.Email)
		return nil
	} else {
		log.Printf("account %s already exists \n", user.Email)
		return fmt.Errorf("account already exists")
	}
}

func Login(credentials shared.CredentialArgs, sessionKey string) error {
	// TODO: NEED CRONJOB PSQL EXTENSION. ANNOYING INSTALL.
	log.Print("Entered db.Login")
	defer log.Print("Exited db.Login")
	rawSQL := `
		WITH user_data AS (
			SELECT user_id
			FROM users
			WHERE email = $1 AND pw = $2
			LIMIT 1
		)
		INSERT INTO user_session (session_key, user_id)
		SELECT $3, user_id
		FROM user_data
	`

	rows, err := db.Exec(
		rawSQL,
		credentials.Email,
		credentials.Password,
		sessionKey,
	)
	if err != nil {
		log.Print(err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Print(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("email and password do not match")
	}

	return nil
}

func Logout(sessionKey string, logoutAllSessions bool) error {
	log.Print("Entered db.Logout")
	defer log.Print("Exited db.Logout")

	rawSQL := ``

	if logoutAllSessions {
		rawSQL = `
		DELETE FROM user_session
		WHERE user_id IN (
			SELECT user_id
			FROM user_sessions
			WHERE session_key = $1
		);`
	} else {
		rawSQL = `DELETE FROM user_session WHERE session_key = $1;`
	}

	rows, err := db.Exec(
		rawSQL,
		sessionKey,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Print(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no such session key")
	} else if rowsAffected > 1 {
		log.Print("all sessions logged out")
	} else if rowsAffected == 1 {
		log.Print("one session logged out")
	}

	return nil
}
