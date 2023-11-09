package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type carRoutes struct {
}

func newCarRoutes() chi.Router {
	routes := &carRoutes{}

	handler := chi.NewRouter()
	handler.Get("/ping", routes.pong)

	return handler
}

func (routes *carRoutes) pong(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
