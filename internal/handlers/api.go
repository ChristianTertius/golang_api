package handlers

import (
	"building_an_api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) // use middleware

	r.Route("/account", func(router chi.Router) {
		// middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
