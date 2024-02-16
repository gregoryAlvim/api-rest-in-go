package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	createBoxRoutes(router)

	return router
}
