package shared

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{
		"error":   errorMsg,
		"message": message,
	}

	json.NewEncoder(w).Encode(response)
}
