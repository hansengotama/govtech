package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/hansengotama/govtech/internal/handler"
	"net/http"
)

func initRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetProducts)
	})

	return r
}
