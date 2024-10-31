package main

import (
	"encoding/json"
	"fmt"
	"go-bbolt/db"
	"go-bbolt/handlers"
	"go-bbolt/models"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func initializeDatabase() (*db.Database, error) {
	// Open the BoltDB
	database, err := db.OpenDB("launch_data.db")
	if err != nil {
		log.Fatalf("Failed to open BoltDB: %v", err)
	}

	// Create the "launches" bucket
	err = database.CreateBucket("launches")
	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	// Read the JSON data
	jsonFile, err := os.Open("../data/launches.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Decode the JSON data into a slice of Launch structs
	var launches []models.Launch
	err = json.NewDecoder(jsonFile).Decode(&launches)
	if err != nil {
		log.Fatalf("Failed to decode JSON data: %v", err)
	}

	// Insert the launch data into the BoltDB
	err = database.InsertLaunchData("launches", launches)
	if err != nil {
		log.Fatalf("Failed to insert data into BoltDB: %v", err)
	}

	fmt.Println("Launch data imported successfully!")

	return database, nil
}

func main() {
	// Initialize database
	database, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Set up the router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Create handler
	h := handlers.NewHandler(database)

	// Setup routes
	router := handlers.SetupRoutes(h)

	// Start server
	port := ":3000"
	log.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}
