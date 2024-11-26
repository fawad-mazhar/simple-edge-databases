# Space Launch Data API (GO LevelDB) 

A RESTful API that manages space launch data using LevelDB as the storage backend. This project provides endpoints to store and retrieve launch information with a focus on performance and simplicity.

## Features

- REST API built with go-chi router
- LevelDB for efficient key-value storage
- Graceful shutdown handling
- Automatic initial data loading

## Prerequisites

- Go 1.21 or higher
- Chi router for HTTP handling
- LevelDB

## Project Structure

```
go-leveldb/
├── README.md
├── api/
│   └── handlers/      # HTTP handlers
│       └── launch_handler.go
├── cmd/
│   └── main.go       # Application entry point
├── models/
│   └── launch.go     # Data models
├── storage/
│   └── leveldb.go    # Database operations
├── utils/
│   └── jsonloader.go # Utility functions
└── go.mod     
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/fawad-mazhar/simple-edge-dbs
cd go-leveldb
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

1. Start the server:
```bash
go run cmd/main.go
```

The server will start on port 8080.

## API Endpoints

### Get All Launches
```
GET /api/launches
```
Returns a list of all launches in the database.

### Get Launch by ID
```
GET /api/launches/{id}
```
Returns a specific launch by its ID.

Example:
```bash
curl http://localhost:8080/api/launches/beb6bfff-c584-42d9-9fea-e3e82a5fd32a
```

## Response Format

The API returns JSON responses in the following format:

```json
{
    "id": "beb6bfff-c584-42d9-9fea-e3e82a5fd32a",
    "name": "Falcon 9 Block 5 | Starlink Group 6-61",
    "status": {
        "id": 2,
        "name": "TBD"
    },
    // ... other launch details
}
```