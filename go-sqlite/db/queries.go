package db

import (
	"database/sql"
	"fmt"

	"go-sqlite/models"
)

func (db *Database) GetAllLaunches() ([]models.Launch, error) {
	launches := []models.Launch{}

	rows, err := db.Query(`SELECT 
		l.id, l.url, l.launch_library_id, l.slug, l.name,
		l.status_id, l.status_name, l.net, l.window_end, l.window_start,
		l.inhold, l.tbdtime, l.tbddate, l.probability, l.holdreason,
		l.failreason, l.hashtag, l.image, l.infographic
	FROM launch l`)
	if err != nil {
		return nil, fmt.Errorf("error querying launches: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var launch models.Launch
		err := rows.Scan(
			&launch.ID, &launch.URL, &launch.LaunchLibraryID, &launch.Slug, &launch.Name,
			&launch.Status.ID, &launch.Status.Name, &launch.Net, &launch.WindowEnd, &launch.WindowStart,
			&launch.Inhold, &launch.TBDtime, &launch.TBDdate, &launch.Probability, &launch.Holdreason,
			&launch.Failreason, &launch.Hashtag, &launch.Image, &launch.Infographic,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning launch: %v", err)
		}

		// Get related data
		if err := db.populateLaunchRelations(&launch); err != nil {
			return nil, err
		}

		launches = append(launches, launch)
	}

	return launches, nil
}

func (db *Database) GetLaunchByID(id string) (*models.Launch, error) {
	var launch models.Launch

	err := db.QueryRow(`SELECT 
		l.id, l.url, l.launch_library_id, l.slug, l.name,
		l.status_id, l.status_name, l.net, l.window_end, l.window_start,
		l.inhold, l.tbdtime, l.tbddate, l.probability, l.holdreason,
		l.failreason, l.hashtag, l.image, l.infographic
	FROM launch l WHERE l.id = ?`, id).Scan(
		&launch.ID, &launch.URL, &launch.LaunchLibraryID, &launch.Slug, &launch.Name,
		&launch.Status.ID, &launch.Status.Name, &launch.Net, &launch.WindowEnd, &launch.WindowStart,
		&launch.Inhold, &launch.TBDtime, &launch.TBDdate, &launch.Probability, &launch.Holdreason,
		&launch.Failreason, &launch.Hashtag, &launch.Image, &launch.Infographic,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying launch: %v", err)
	}

	// Get related data
	if err := db.populateLaunchRelations(&launch); err != nil {
		return nil, err
	}

	return &launch, nil
}

func (db *Database) populateLaunchRelations(launch *models.Launch) error {
	// Get launch service provider
	err := db.QueryRow(`SELECT id, url, name, type FROM launch_service_provider WHERE id = ?`,
		launch.LaunchServiceProvider.ID).Scan(
		&launch.LaunchServiceProvider.ID,
		&launch.LaunchServiceProvider.URL,
		&launch.LaunchServiceProvider.Name,
		&launch.LaunchServiceProvider.Type,
	)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying launch service provider: %v", err)
	}

	// Get rocket and configuration
	err = db.QueryRow(`SELECT 
		id, configuration_id, configuration_launch_library_id,
		configuration_url, configuration_name, configuration_family,
		configuration_full_name, configuration_variant
	FROM rocket WHERE id = ?`, launch.Rocket.ID).Scan(
		&launch.Rocket.ID,
		&launch.Rocket.Configuration.ID,
		&launch.Rocket.Configuration.LaunchLibraryID,
		&launch.Rocket.Configuration.URL,
		&launch.Rocket.Configuration.Name,
		&launch.Rocket.Configuration.Family,
		&launch.Rocket.Configuration.FullName,
		&launch.Rocket.Configuration.Variant,
	)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying rocket: %v", err)
	}

	// Get mission
	err = db.QueryRow(`SELECT 
		id, launch_library_id, name, description, launch_designator,
		type, orbit_id, orbit_name, orbit_abbrev
	FROM mission WHERE id = ?`, launch.Mission.ID).Scan(
		&launch.Mission.ID,
		&launch.Mission.LaunchLibraryID,
		&launch.Mission.Name,
		&launch.Mission.Description,
		&launch.Mission.LaunchDesignator,
		&launch.Mission.Type,
		&launch.Mission.Orbit.ID,
		&launch.Mission.Orbit.Name,
		&launch.Mission.Orbit.Abbrev,
	)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying mission: %v", err)
	}

	// Get pad and location
	err = db.QueryRow(`SELECT 
		id, url, agency_id, name, info_url, wiki_url, map_url,
		latitude, longitude, location_id, location_url, location_name,
		location_country_code, location_map_image, location_total_launch_count,
		location_total_landing_count, map_image, total_launch_count
	FROM pad WHERE id = ?`, launch.Pad.ID).Scan(
		&launch.Pad.ID,
		&launch.Pad.URL,
		&launch.Pad.AgencyID,
		&launch.Pad.Name,
		&launch.Pad.InfoURL,
		&launch.Pad.WikiURL,
		&launch.Pad.MapURL,
		&launch.Pad.Latitude,
		&launch.Pad.Longitude,
		&launch.Pad.Location.ID,
		&launch.Pad.Location.URL,
		&launch.Pad.Location.Name,
		&launch.Pad.Location.CountryCode,
		&launch.Pad.Location.MapImage,
		&launch.Pad.Location.TotalLaunchCount,
		&launch.Pad.Location.TotalLandingCount,
		&launch.Pad.MapImage,
		&launch.Pad.TotalLaunchCount,
	)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying pad: %v", err)
	}

	// Get programs
	rows, err := db.Query(`
		SELECT p.id, p.url, p.name, p.description, p.image_url,
			p.start_date, p.end_date, p.info_url, p.wiki_url
		FROM program p
		JOIN launch_program lp ON p.id = lp.program_id
		WHERE lp.launch_id = ?`, launch.ID)
	if err != nil {
		return fmt.Errorf("error querying programs: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var program models.Program
		err := rows.Scan(
			&program.ID,
			&program.URL,
			&program.Name,
			&program.Description,
			&program.ImageURL,
			&program.StartDate,
			&program.EndDate,
			&program.InfoURL,
			&program.WikiURL,
		)
		if err != nil {
			return fmt.Errorf("error scanning program: %v", err)
		}

		// Get program agencies
		agencyRows, err := db.Query(`
			SELECT agency_id, agency_url, agency_name, agency_type
			FROM program_agency
			WHERE program_id = ?`, program.ID)
		if err != nil {
			return fmt.Errorf("error querying program agencies: %v", err)
		}
		defer agencyRows.Close()

		for agencyRows.Next() {
			var agency models.Agency
			err := agencyRows.Scan(
				&agency.ID,
				&agency.URL,
				&agency.Name,
				&agency.Type,
			)
			if err != nil {
				return fmt.Errorf("error scanning agency: %v", err)
			}
			program.Agencies = append(program.Agencies, agency)
		}

		launch.Program = append(launch.Program, program)
	}

	return nil
}
