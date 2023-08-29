package controllers

import (
	"encoding/json"
	"leslies-app/backend/services"
	"leslies-app/backend/shared"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Server probably OK")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered controllers.CreateUser")
	defer log.Println("Exited controllers.CreateUser")

	// need to validate input

	var user shared.CreateUserArgs

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse JSON request body.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = ValidateCreateUserArgs(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.FirstName = strings.ToUpper(user.FirstName[:1]) + user.FirstName[1:]
	user.LastName = strings.ToUpper(user.LastName[:1]) + user.LastName[1:]

	err = services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter controllers.fmtin")
	defer log.Println("Exited controllers.fmtin")

	var credentials shared.CredentialArgs

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Failed to parse JSON request body.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	sessionKey, err := services.Login(credentials)
	if err != nil {
		http.Error(w, "Internal error.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", sessionKey)
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter controllers.fmtin")
	defer log.Println("Exited controllers.fmtin")

	sessionKey := r.Header.Get("Authorization")
	logoutAllSessions := r.Header.Get("LogoutAllSessions")

	logoutAllSessionsBool, err := strconv.ParseBool(logoutAllSessions)

	if err != nil {
		http.Error(w, "Check logOutAllSessions bool.", http.StatusInternalServerError)
		return
	}

	err = services.Logout(sessionKey, logoutAllSessionsBool)

	if err != nil {
		http.Error(w, "Unable to logout.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func DeleteAccount(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Enter controllers.fmtin")
// 	defer log.Println("Exited controllers.fmtin")

// 	var credentials shared.CredentialArgs

// 	err := json.NewDecoder(r.Body).Decode(&credentials)
// 	if err != nil {
// 		http.Error(w, "Failed to parse JSON request body.", http.StatusBadRequest)
// 		return
// 	}

// 	defer r.Body.Close()

// 	sessionKey, err := services.Login(credentials)
// 	if err != nil {
// 		http.Error(w, "Internal error.", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Authorization", sessionKey)
// 	w.WriteHeader(http.StatusOK)
// }
