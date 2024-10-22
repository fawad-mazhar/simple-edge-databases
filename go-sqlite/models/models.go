package models

type Launch struct {
	ID                    string    `json:"id"`
	URL                   string    `json:"url"`
	LaunchLibraryID       *int      `json:"launch_library_id"`
	Slug                  string    `json:"slug"`
	Name                  string    `json:"name"`
	Status                Status    `json:"status"`
	Net                   string    `json:"net"`
	WindowEnd             string    `json:"window_end"`
	WindowStart           string    `json:"window_start"`
	Inhold                bool      `json:"inhold"`
	TBDtime               bool      `json:"tbdtime"`
	TBDdate               bool      `json:"tbddate"`
	Probability           *int      `json:"probability"`
	Holdreason            string    `json:"holdreason"`
	Failreason            string    `json:"failreason"`
	Hashtag               *string   `json:"hashtag"`
	LaunchServiceProvider Provider  `json:"launch_service_provider"`
	Rocket                Rocket    `json:"rocket"`
	Mission               Mission   `json:"mission"`
	Pad                   Pad       `json:"pad"`
	WebcastLive           bool      `json:"webcast_live"`
	Image                 string    `json:"image"`
	Infographic           *string   `json:"infographic"`
	Program               []Program `json:"program"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Provider struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Rocket struct {
	ID            int           `json:"id"`
	Configuration Configuration `json:"configuration"`
}

type Configuration struct {
	ID              int    `json:"id"`
	LaunchLibraryID *int   `json:"launch_library_id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	Family          string `json:"family"`
	FullName        string `json:"full_name"`
	Variant         string `json:"variant"`
}

type Mission struct {
	ID               int     `json:"id"`
	LaunchLibraryID  *int    `json:"launch_library_id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	LaunchDesignator *string `json:"launch_designator"`
	Type             string  `json:"type"`
	Orbit            Orbit   `json:"orbit"`
}

type Orbit struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
}

type Pad struct {
	ID               int      `json:"id"`
	URL              string   `json:"url"`
	AgencyID         *int     `json:"agency_id"`
	Name             string   `json:"name"`
	InfoURL          *string  `json:"info_url"`
	WikiURL          *string  `json:"wiki_url"`
	MapURL           string   `json:"map_url"`
	Latitude         string   `json:"latitude"`
	Longitude        string   `json:"longitude"`
	Location         Location `json:"location"`
	MapImage         string   `json:"map_image"`
	TotalLaunchCount int      `json:"total_launch_count"`
}

type Location struct {
	ID                int    `json:"id"`
	URL               string `json:"url"`
	Name              string `json:"name"`
	CountryCode       string `json:"country_code"`
	MapImage          string `json:"map_image"`
	TotalLaunchCount  int    `json:"total_launch_count"`
	TotalLandingCount int    `json:"total_landing_count"`
}

type Program struct {
	ID          int      `json:"id"`
	URL         string   `json:"url"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Agencies    []Agency `json:"agencies"`
	ImageURL    string   `json:"image_url"`
	StartDate   string   `json:"start_date"`
	EndDate     *string  `json:"end_date"`
	InfoURL     string   `json:"info_url"`
	WikiURL     string   `json:"wiki_url"`
}

type Agency struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
}
