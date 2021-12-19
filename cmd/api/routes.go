package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/v8tix/kit/handler"
	"net/http"
)

func Routes(handler handler.Handler) http.Handler {
	r := chi.NewRouter()
	r.NotFound(handler.NotFoundResponse)
	r.MethodNotAllowed(handler.MethodNotAllowedResponse)

	r.Route("/public", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handler.RecoverPanic)
			r.Use(handler.EnableCORS)
			r.Get("/health", handler.HealthCheck)
		})
	})

	/*	r.Route("/beers", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handler.RecoverPanic)
			r.Use(handler.EnableCORS)
			r.Post("/", handler.PostBeer)
		})
	})*/

	return r
}
