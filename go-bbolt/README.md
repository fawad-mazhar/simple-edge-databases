# Space Launch Data API (Bolt DB)

A RESTful API service that manages and provides information about space launches. The service stores launch data in Bolt DB and provides endpoints to query launch information.

## Features

- BoltDB (included as a dependency)
- RESTful API endpoints
- Automatic database initialization and data loading
- Detailed launch information including rockets, missions, and launch sites

## Tech Stack

- Go 1.21+
- Chi router for HTTP handling
- SQLite3 for data storage

## Project Structure

```
.
├── README.md
├── db
│   ├── db.go              # Database connection and operations
│   └── queries.go         # Database queries
├── go.mod
├── go.sum
├── handlers
│   ├── handlers.go        # HTTP handlers
│   └── routes.go          # Route definitions
├── internal
│   └── utils
│       └── utils.go
├── launch_data.db
├── main.go                # Application entry point
├── models                 # Data structures
│   ├── bucket.go
│   └── models.go
└── ui                     # BoltDB Browser
    ├── static
    │   ├── script.js
    │   └── style.css
    └── templates
        ├── bucket.html
        └── buckets.html

```

## Prerequisites

- Go 1.21 or higher
- Bbolt

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fawad1985/sample-edge-dbs
   cd go-bbolt
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```