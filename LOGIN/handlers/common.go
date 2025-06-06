package handlers

import (
	"encoding/json"
	"net/http"

	"Login/LOGIN/models"
)

// RespondWithJSON envía una respuesta JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError envía una respuesta de error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, models.Response{
		Success: false,
		Message: message,
	})
}