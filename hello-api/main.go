package main

import (
	"log"
	"net/http"

	"hello-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.Hello)

	log.Fatal(http.ListenAndServe(":9090", mux))
}
