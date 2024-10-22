package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Routes
	r.Route("/api", func(r chi.Router) {
		// Health check endpoint
		r.Get("/health", h.HealthCheck)

		// Launches endpoints
		r.Route("/launches", func(r chi.Router) {
			r.Get("/", h.GetAllLaunches)
			r.Get("/{id}", h.GetLaunchByID)
		})
	})

	return r
}
