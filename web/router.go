package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebSettings struct {
	Port string
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return r
}

func SetupRoutes(r *chi.Mux, h *Handlers) {
	r.Get("/", h.GetIndex)
}

func ListenAndServe(settings *WebSettings, router *chi.Mux) error {
	return http.ListenAndServe(settings.Port, router)
}
