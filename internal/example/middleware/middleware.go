package middleware

import (
	"net/http"
	"encoding/json"
	"go_example/internal/example/models"
)

func Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response {
		Message: "test message",
	}

	json.NewEncoder(w).Encode(response)
}

func GetAnon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string
	}{
		Message: "test with anonymous struct",
	}

	json.NewEncoder(w).Encode(response)
}