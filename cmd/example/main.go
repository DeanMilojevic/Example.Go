package main

import (
	"go_example/internal/example/router"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
