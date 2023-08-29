package middleware

import (
	_ "crypto/sha256"
	"encoding/json"
	"leslies-app/backend/db"
	"log"
	"net/http"
)

func unauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	response := map[string]string{
		"error":   "Unauthorized",
		"message": message,
	}

	json.NewEncoder(w).Encode(response)
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Enter middleware.Auth")
		defer log.Println("Exited middleware.Auth")
		authenticated := false

		sessionKey := r.Header.Get("Authorization")

		authenticated, err := db.Auth(sessionKey)
		if err != nil {
			http.Error(w, "Internal error.", http.StatusInternalServerError)
			return
		}

		if authenticated {
			next.ServeHTTP(w, r)
		} else {
			unauthorizedResponse(w, "Authentication failed. Please provide valid credentials.")
			return
		}
	}
}
