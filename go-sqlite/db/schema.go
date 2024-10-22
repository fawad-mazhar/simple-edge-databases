package db

import (
	"fmt"
)

func (db *Database) InitSchema() error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	defer tx.Rollback()

	// Array of table creation statements
	schemas := []string{
		`CREATE TABLE IF NOT EXISTS launch (
			id TEXT PRIMARY KEY,
			url TEXT,
			launch_library_id TEXT,
			slug TEXT,
			name TEXT,
			status_id INTEGER,
			status_name TEXT,
			net TEXT,
			window_end TEXT,
			window_start TEXT,
			inhold BOOLEAN,
			tbdtime BOOLEAN,
			tbddate BOOLEAN,
			probability INTEGER,
			holdreason TEXT,
			failreason TEXT,
			hashtag TEXT,
			image TEXT,
			infographic TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS launch_service_provider (
			id INTEGER PRIMARY KEY,
			url TEXT,
			name TEXT,
			type TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS rocket (
			id INTEGER PRIMARY KEY,
			configuration_id INTEGER,
			configuration_launch_library_id INTEGER,
			configuration_url TEXT,
			configuration_name TEXT,
			configuration_family TEXT,
			configuration_full_name TEXT,
			configuration_variant TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS mission (
			id INTEGER PRIMARY KEY,
			launch_library_id INTEGER,
			name TEXT,
			description TEXT,
			launch_designator TEXT,
			type TEXT,
			orbit_id INTEGER,
			orbit_name TEXT,
			orbit_abbrev TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS pad (
			id INTEGER PRIMARY KEY,
			url TEXT,
			agency_id INTEGER,
			name TEXT,
			info_url TEXT,
			wiki_url TEXT,
			map_url TEXT,
			latitude TEXT,
			longitude TEXT,
			location_id INTEGER,
			location_url TEXT,
			location_name TEXT,
			location_country_code TEXT,
			location_map_image TEXT,
			location_total_launch_count INTEGER,
			location_total_landing_count INTEGER,
			map_image TEXT,
			total_launch_count INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS program (
			id INTEGER PRIMARY KEY,
			url TEXT,
			name TEXT,
			description TEXT,
			image_url TEXT,
			start_date TEXT,
			end_date TEXT,
			info_url TEXT,
			wiki_url TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS launch_program (
			launch_id TEXT,
			program_id INTEGER,
			PRIMARY KEY (launch_id, program_id),
			FOREIGN KEY (launch_id) REFERENCES launch(id),
			FOREIGN KEY (program_id) REFERENCES program(id)
		)`,
		`CREATE TABLE IF NOT EXISTS program_agency (
			program_id INTEGER,
			agency_id INTEGER,
			agency_url TEXT,
			agency_name TEXT,
			agency_type TEXT,
			PRIMARY KEY (program_id, agency_id),
			FOREIGN KEY (program_id) REFERENCES program(id)
		)`,
	}

	// Execute each CREATE TABLE statement
	for _, schema := range schemas {
		if _, err := tx.Exec(schema); err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing schema transaction: %v", err)
	}

	return nil
}
