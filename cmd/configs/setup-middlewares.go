package configs

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupMiddlewares(router *chi.Mux) {
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}
