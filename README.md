# Simple Edge Databases
---

Fawad Mazhar <fawadmazhar@hotmail.com> 2024

---

# Simple Edge Databases

This repository demonstrates implementations of three different edge/embedded databases in Go, each powering a REST API for managing space launch data. The project showcases how to work with various embedded databases that are suitable for edge computing scenarios.

## Overview

The repository contains three independent implementations using different database technologies:

1. **BBolt Implementation** (`/go-bbolt`)
   - Uses BBolt (formerly Bolt DB), a pure Go key/value store
   - Includes a web-based database browser UI
   - Perfect for scenarios requiring ACID compliance with a simple key-value model

2. **LevelDB Implementation** (`/go-leveldb`)
   - Uses Google's LevelDB, an efficient key-value storage library
   - Optimized for high-performance reads and writes
   - Ideal for applications requiring fast access to large datasets

3. **SQLite Implementation** (`/go-sqlite`)
   - Uses SQLite, a reliable embedded relational database
   - Provides full SQL query capabilities
   - Best suited for applications needing structured data and complex queries

Each implementation provides the same REST API interface for managing space launch data, allowing for easy comparison between the different database solutions.

## Project Structure

```
simple-edge-databases/
├── data/
│   └── launches.json          # Shared launch data
├── go-bbolt/                  # BBolt implementation
├── go-leveldb/               # LevelDB implementation
└── go-sqlite/                # SQLite implementation
```

## Common Features Across Implementations

- RESTful API endpoints for managing launch data
- Automatic database initialization
- Initial data loading from shared JSON file
- Similar API response formats
- Built with Go 1.21+
- Uses Chi router for HTTP handling

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/fawad-mazhar/simple-edge-dbs
cd simple-edge-dbs
```

2. Choose an implementation:
   - For BBolt: `cd go-bbolt`
   - For LevelDB: `cd go-leveldb`
   - For SQLite: `cd go-sqlite`

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run main.go    # For BBolt and SQLite
go run cmd/main.go # For LevelDB
```

## API Endpoints

All implementations provide similar endpoints:

```
GET /api/launches      # Get all launches
GET /api/launches/{id} # Get launch by ID
```

# Simple Edge Databases

[Previous sections remain the same...]

## Choosing the Right Implementation

- **Choose BBolt if you need:**
  - ACID compliance with all-or-nothing transactions
  - Simple key-value storage with bucket-based organization (similar to collections/tables)
  - Single-file database that's easy to backup and restore
  - Pure Go implementation with zero external dependencies
  - Concurrent read access with exclusive write locks
  - Memory-mapped file access for better performance
  - Built-in database browser UI for easy data inspection
  - Perfect for embedded systems and applications requiring data integrity
  - Best suited for workloads where writes are less frequent than reads

- **Choose LevelDB if you need:**
  - High-performance reads and writes using Log-Structured Merge (LSM) trees
  - Efficient storage for large datasets with automatic compression
  - Ordered key-value pairs enabling range queries
  - Minimal memory overhead with configurable cache sizes
  - Multi-threaded compaction for better write performance
  - Snapshot isolation for consistent reads
  - Ideal for high-throughput applications and time-series data
  - Great for write-heavy workloads and sequential reads
  - Note: No ACID guarantees, but provides atomic batch operations

- **Choose SQLite if you need:**
  - Full SQL query capabilities with JOIN operations
  - Complex data relationships and foreign key constraints
  - Structured data with schema enforcement
  - Robust indexing capabilities for optimized queries
  - Transaction support with ACID properties
  - Zero-configuration required to get started
  - Single-file database with cross-platform compatibility
  - Rich ecosystem of tools and viewers
  - Ideal for applications requiring structured queries and complex data relationships
  - Best for traditional CRUD applications with relational data models

### Performance Considerations

- **BBolt**
  - Optimized for read operations
  - Write operations acquire a file lock, limiting concurrent writes
  - Memory-mapped files provide fast read access
  - Scales well with database size due to B+ tree structure

- **LevelDB**
  - Excellent write performance due to LSM tree structure
  - Good read performance for recently written data
  - Compaction process may impact performance periodically
  - Scales well horizontally with data size

- **SQLite**
  - Balanced read/write performance
  - Excellent for complex queries and joins
  - Performance can be optimized with proper indexing
  - May require more careful query optimization for large datasets

### Storage Patterns

- **BBolt**
  ```
  Bucket (Collection)
  └── Key-Value Pairs
      └── Nested Buckets (Optional)
  ```

- **LevelDB**
  ```
  Key-Value Store
  └── Sorted Key-Value Pairs
      └── No nested structures
  ```

- **SQLite**
  ```
  Database
  └── Tables
      └── Rows and Columns
          └── Foreign Key Relationships
  ```

[Rest of the README remains the same...]

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Detailed Documentation

For more detailed information about each implementation, please refer to their respective README files:

- [BBolt Implementation](go-bbolt/README.md)
- [LevelDB Implementation](go-leveldb/README.md)
- [SQLite Implementation](go-sqlite/README.md)