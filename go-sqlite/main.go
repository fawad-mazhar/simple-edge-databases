package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"go-sqlite/db"
	"go-sqlite/handlers"
	"go-sqlite/models"
)

type Server struct {
	db *db.Database
}

// Initialize database and load initial data
func initializeDatabase() (*db.Database, error) {
	// Connect to SQLite database
	database, err := db.NewDatabase("./launches.db")
	if err != nil {
		return nil, err
	}

	// Initialize schema
	log.Println("Initializing database schema...")
	if err := database.InitSchema(); err != nil {
		return nil, err
	}
	log.Println("Schema initialization completed")

	// Check if we already have data
	launches, err := database.GetAllLaunches()
	if err != nil {
		return nil, err
	}

	// If we already have data, skip initialization
	if len(launches) > 0 {
		log.Printf("Database already contains %d launches, skipping data initialization\n", len(launches))
		return database, nil
	}

	// Load initial data from JSON file
	log.Println("Loading initial data from launches.json...")
	jsonFile, err := os.Open("../data/launches.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var launchData []models.Launch
	if err := json.Unmarshal(byteValue, &launchData); err != nil {
		return nil, err
	}

	// Insert launches
	log.Printf("Inserting %d launches into database...\n", len(launchData))
	for i, launch := range launchData {
		if err := database.InsertLaunchTransaction(launch); err != nil {
			log.Printf("Warning: Error inserting launch %s: %v\n", launch.ID, err)
		} else {
			log.Printf("Inserted launch %d/%d: %s\n", i+1, len(launchData), launch.ID)
		}
	}
	log.Println("Data initialization completed")

	return database, nil
}

func main() {
	// Initialize database
	database, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

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
