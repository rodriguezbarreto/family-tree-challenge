package main

import (
	"famlily-tree-challenge/cmd/configs"
	"famlily-tree-challenge/cmd/enpoints"

	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	configs.SetupMiddlewares(router)
	enpoints.SetupEndpoints(router)

	http.ListenAndServe(":8080", router)
}
