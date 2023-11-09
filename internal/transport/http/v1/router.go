package v1

import "github.com/go-chi/chi/v5"

func NewRouter(handler chi.Router) {
	handler.Route("/api/v1", func(r chi.Router) {
		r.Mount("/car", newCarRoutes())
	})
}
