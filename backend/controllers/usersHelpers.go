package controllers

import (
	"fmt"
	"leslies-app/backend/shared"
	"regexp"
)

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func isValidPassword(password string) bool {
	return len(password) > 0
}

func ValidateCreateUserArgs(args shared.CreateUserArgs) error {
	if args.FirstName == "" {
		return fmt.Errorf("first name must be non-empty")
	}

	if args.LastName == "" {
		return fmt.Errorf("last name must be non-empty")
	}

	if !isValidEmail(args.Email) {
		return fmt.Errorf("invalid email address")
	}

	if !isValidPassword(args.Password) {
		return fmt.Errorf("invalid password")
	}

	return nil
}
