package main

import (
	"github.com/go-chi/chi/v5"
	. "github.com/v8tix/bender/internal/handler"
	"net/http"
)

func Routes(benderHandler BenderHandler) http.Handler {
	r := chi.NewRouter()
	r.NotFound(benderHandler.NotFoundResponse)
	r.MethodNotAllowed(benderHandler.MethodNotAllowedResponse)

	r.Route("/public", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(benderHandler.RecoverPanic)
			r.Use(benderHandler.EnableCORS)
			r.Get("/health", benderHandler.HealthCheck)
			r.Post("/task", benderHandler.ExecBackgroundTask)
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
