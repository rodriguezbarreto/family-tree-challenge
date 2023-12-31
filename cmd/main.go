package main

import (
	"family-tree-challenge/cmd/configs"
	"family-tree-challenge/cmd/enpoints"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()

	configs.SetupMiddlewares(router)
	enpoints.SetupEndpoints(router)

	port := ":8080"
	log.Printf("Server is starting on port %s...", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
