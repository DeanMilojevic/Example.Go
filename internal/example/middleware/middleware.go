package middleware

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
}

func Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("ContentType", "application/json")

	response := Response {
		Message: "test message",
	}

	json.NewEncoder(w).Encode(response)
}