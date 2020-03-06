package middleware

import (
	"net/http"
	"encoding/json"
	"go_example/internal/example/models"
)

func Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("ContentType", "application/json")

	response := models.Response {
		Message: "test message",
	}

	json.NewEncoder(w).Encode(response)
}