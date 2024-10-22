# Space Launches API (GO SQLITE)

A RESTful API service that manages and provides information about space launches. The service stores launch data in SQLite and provides endpoints to query launch information.

## Features

- SQLite database for storing launch data
- RESTful API endpoints
- Automatic database initialization and data loading
- Detailed launch information including rockets, missions, and launch sites

## Tech Stack

- Go 1.21+
- Chi router for HTTP handling
- SQLite3 for data storage
- Standard library's encoding/json for JSON processing

## Project Structure

```
.
├── README.md
├── main.go                 # Application entry point
├── models/
│   └── models.go          # Data structures
├── db/
│   ├── db.go             # Database connection and operations
│   ├── schema.go         # Database schema
│   └── queries.go        # Database queries
├── handlers/
│   ├── handlers.go       # HTTP handlers
│   └── routes.go         # Route definitions
```

## Prerequisites

- Go 1.21 or higher
- SQLite3

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fawad1985/sample-edge-dbs
   cd go-sqlite
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

   The server will:
   - Initialize the SQLite database
   - Create necessary tables
   - Load initial data from launches.json (if the database is empty)
   - Start the HTTP server on port 3000

## API Endpoints

### Health Check
```
GET /api/health
```
Returns the service health status.

Response:
```json
{
    "status": "ok"
}
```

### Get All Launches
```
GET /api/launches
```
Returns an array of all launches in the database.

Response:
```json
[
    {
        "id": "beb6bfff-c584-42d9-9fea-e3e82a5fd32a",
        "url": "https://example.com/launch/1",
        "name": "Falcon 9 Block 5 | Starlink Group 6-61",
        "status": {
            "id": 2,
            "name": "TBD"
        },
        // ... other launch details
    },
    // ... more launches
]
```

### Get Launch by ID
```
GET /api/launches/{id}
```
Returns details for a specific launch.

Response:
```json
{
    "id": "beb6bfff-c584-42d9-9fea-e3e82a5fd32a",
    "url": "https://example.com/launch/1",
    "name": "Falcon 9 Block 5 | Starlink Group 6-61",
    "status": {
        "id": 2,
        "name": "TBD"
    },
    "mission": {
        "name": "Starlink Group 6-61",
        "description": "A batch of Starlink satellites",
        // ... other mission details
    },
    // ... other launch details
}
```

## Error Responses

The API uses standard HTTP status codes:

- `200 OK` - Success
- `404 Not Found` - Launch not found
- `500 Internal Server Error` - Server error

Error response format:
```json
{
    "error": "Error message here"
}
```