package handlers

import (
	"net/http"
	"path/filepath"

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

	workDir, _ := filepath.Abs(".")
	filesDir := http.Dir(filepath.Join(workDir, "ui/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(filesDir)))

	// Routes
	r.Get("/", h.Buckets)
	r.Get("/bucket/{name}", h.Bucket)

	r.Route("/api", func(r chi.Router) {
		// Health check endpoint
		r.Get("/health", h.HealthCheck)

		// Launches endpoints
		r.Route("/launches", func(r chi.Router) {
			r.Get("/", h.GetAllLaunches)
			r.Get("/{id}", h.GetLaunchByID)
		})

		r.Get("/bucket/{bucket}/key/{key}", h.KeyValue)
	})

	return r
}
