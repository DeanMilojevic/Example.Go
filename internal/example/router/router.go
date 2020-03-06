package router

import (
	"github.com/gorilla/mux"

	"go_example/internal/example/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/example", middleware.GetAnon).Methods("GET", "OPTIONS")

	return router
}
