package web

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	return r
}

func SetupRoutes(r *chi.Mux, h *Handlers) {
	r.Get("/", h.GetIndex)
}

func ListenAndServe(settings *WebSettings, router *chi.Mux) error {
	return http.ListenAndServe(settings.Port, router)
}
