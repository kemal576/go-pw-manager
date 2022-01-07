package api

import (
	"encoding/json"
	"net/http"

	"github.com/kemal576/go-pw-manager/models"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, &models.Error{Status: "Error", Message: message})
}
