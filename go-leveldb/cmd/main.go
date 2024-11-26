package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-leveldb/api/handlers"
	"go-leveldb/storage"
	"go-leveldb/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize storage
	store, err := storage.NewLaunchStore("launches.db")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	// Load initial data if database is empty
	if err := loadInitialData(store); err != nil {
		log.Fatal(err)
	}

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize handlers
	launchHandler := handlers.NewLaunchHandler(store)

	// Routes
	r.Route("/api/launches", func(r chi.Router) {
		r.Get("/", launchHandler.GetAllLaunches)
		r.Get("/{id}", launchHandler.GetLaunchByID)
	})

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Println("Server started on :8080")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}

func loadInitialData(store *storage.LaunchStore) error {
	// Check if we already have data
	launches, err := store.GetAllLaunches()
	if err != nil {
		return err
	}

	// If we have data, skip loading
	if len(launches) > 0 {
		return nil
	}

	// Load initial data from JSON
	launches, err = utils.LoadLaunchesFromJSON("../data/launches.json")
	if err != nil {
		return err
	}

	// Store launches in database
	return store.StoreLaunches(launches)
}
