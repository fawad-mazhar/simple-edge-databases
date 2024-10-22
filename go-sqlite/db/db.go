package db

import (
	"database/sql"
	"fmt"

	"go-sqlite/models"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) InsertLaunchTransaction(launch models.Launch) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	defer tx.Rollback() // Rollback the transaction if it hasn't been committed

	// Insert launch
	if err := db.insertLaunch(tx, launch); err != nil {
		return fmt.Errorf("error inserting launch: %v", err)
	}

	// Insert launch service provider
	if err := db.insertLaunchServiceProvider(tx, launch.LaunchServiceProvider); err != nil {
		return fmt.Errorf("error inserting launch service provider: %v", err)
	}

	// Insert rocket
	if err := db.insertRocket(tx, launch.Rocket); err != nil {
		return fmt.Errorf("error inserting rocket: %v", err)
	}

	// Insert mission
	if err := db.insertMission(tx, launch.Mission); err != nil {
		return fmt.Errorf("error inserting mission: %v", err)
	}

	// Insert pad
	if err := db.insertPad(tx, launch.Pad); err != nil {
		return fmt.Errorf("error inserting pad: %v", err)
	}

	// Insert programs and program agencies
	for _, program := range launch.Program {
		if err := db.insertProgram(tx, program); err != nil {
			return fmt.Errorf("error inserting program: %v", err)
		}

		if err := db.insertLaunchProgram(tx, launch.ID, program.ID); err != nil {
			return fmt.Errorf("error inserting launch_program: %v", err)
		}

		for _, agency := range program.Agencies {
			if err := db.insertProgramAgency(tx, program.ID, agency); err != nil {
				return fmt.Errorf("error inserting program_agency: %v", err)
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (db *Database) insertLaunch(tx *sql.Tx, launch models.Launch) error {
	_, err := tx.Exec(`INSERT INTO launch (id, url, launch_library_id, slug, name, status_id, status_name, net, window_end, window_start, inhold, tbdtime, tbddate, probability, holdreason, failreason, hashtag, image, infographic) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		launch.ID, launch.URL, launch.LaunchLibraryID, launch.Slug, launch.Name, launch.Status.ID, launch.Status.Name, launch.Net, launch.WindowEnd, launch.WindowStart, launch.Inhold, launch.TBDtime, launch.TBDdate, launch.Probability, launch.Holdreason, launch.Failreason, launch.Hashtag, launch.Image, launch.Infographic)
	return err
}

func (db *Database) insertLaunchServiceProvider(tx *sql.Tx, provider models.Provider) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO launch_service_provider (id, url, name, type) VALUES (?, ?, ?, ?)`,
		provider.ID, provider.URL, provider.Name, provider.Type)
	return err
}

func (db *Database) insertRocket(tx *sql.Tx, rocket models.Rocket) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO rocket (id, configuration_id, configuration_launch_library_id, configuration_url, configuration_name, configuration_family, configuration_full_name, configuration_variant) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		rocket.ID, rocket.Configuration.ID, rocket.Configuration.LaunchLibraryID, rocket.Configuration.URL, rocket.Configuration.Name, rocket.Configuration.Family, rocket.Configuration.FullName, rocket.Configuration.Variant)
	return err
}

func (db *Database) insertMission(tx *sql.Tx, mission models.Mission) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO mission (id, launch_library_id, name, description, launch_designator, type, orbit_id, orbit_name, orbit_abbrev) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		mission.ID, mission.LaunchLibraryID, mission.Name, mission.Description, mission.LaunchDesignator, mission.Type, mission.Orbit.ID, mission.Orbit.Name, mission.Orbit.Abbrev)
	return err
}

func (db *Database) insertPad(tx *sql.Tx, pad models.Pad) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO pad (id, url, agency_id, name, info_url, wiki_url, map_url, latitude, longitude, location_id, location_url, location_name, location_country_code, location_map_image, location_total_launch_count, location_total_landing_count, map_image, total_launch_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		pad.ID, pad.URL, pad.AgencyID, pad.Name, pad.InfoURL, pad.WikiURL, pad.MapURL, pad.Latitude, pad.Longitude, pad.Location.ID, pad.Location.URL, pad.Location.Name, pad.Location.CountryCode, pad.Location.MapImage, pad.Location.TotalLaunchCount, pad.Location.TotalLandingCount, pad.MapImage, pad.TotalLaunchCount)
	return err
}

func (db *Database) insertProgram(tx *sql.Tx, program models.Program) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO program (id, url, name, description, image_url, start_date, end_date, info_url, wiki_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		program.ID, program.URL, program.Name, program.Description, program.ImageURL, program.StartDate, program.EndDate, program.InfoURL, program.WikiURL)
	return err
}

func (db *Database) insertLaunchProgram(tx *sql.Tx, launchID string, programID int) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO launch_program (launch_id, program_id) VALUES (?, ?)`,
		launchID, programID)
	return err
}

func (db *Database) insertProgramAgency(tx *sql.Tx, programID int, agency models.Agency) error {
	_, err := tx.Exec(`INSERT OR REPLACE INTO program_agency (program_id, agency_id, agency_url, agency_name, agency_type) VALUES (?, ?, ?, ?, ?)`,
		programID, agency.ID, agency.URL, agency.Name, agency.Type)
	return err
}
