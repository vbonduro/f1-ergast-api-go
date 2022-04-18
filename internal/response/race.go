package response

// TimeGroup has schedule information for when an event occurs.
type TimeGroup struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// Location has location information for a RaceCircuit in a race schedule JSON response.
type Location struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"long"`
	Locality  string `json:"locality"`
	Country   string `json:"country"`
}

//RaceCircuit contains circuit information from the race schedule JSON response.
type RaceCircuit struct {
	Id           string   `json:"circuitId"`
	Url          string   `json:"url"`
	CircutName   string   `json:"circuitName"`
	LocationInfo Location `json:"Location"`
}

// RaceSchedule contains the JSON response structure from a Get request
// to the race schedule URI.
type RaceSchedule struct {
	Season         string      `json:"season"`
	Round          string      `json:"round"`
	Url            string      `json:"url"`
	Name           string      `json:"raceName"`
	Circuit        RaceCircuit `json:"Circuit"`
	Date           string      `json:"date"`
	Time           string      `json:"time"`
	FirstPractice  TimeGroup   `json:"FirstPractice"`
	Sprint         TimeGroup   `json:"Sprint"`
	SecondPractice TimeGroup   `json:"SecondPractice"`
	ThirdPractice  TimeGroup   `json:"ThirdPractice"`
	Qualifying     TimeGroup   `json:"Qualifying"`
}

// RaceTable contains the actual race schedule for a season.
type RaceTable struct {
	Season string         `json:"season"`
	Races  []RaceSchedule `json:"Races"`
}

// ScheduleHeader has header information for the race schedule.
type ScheduleHeader struct {
	// todo: skipping a bunch of stuff.
	ScheduleTable RaceTable `json:"RaceTable"`
}

// ScheduleResponse is the top of the response tree for a Get request for
// the race schedule.
type ScheduleResponse struct {
	Header ScheduleHeader `json:"MRData"`
}
