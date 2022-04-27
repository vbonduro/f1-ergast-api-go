package ergast

import (
	"github.com/vbonduro/f1-ergast-api-go/internal/request"
)

// Constructor has constructor information from a driver standings JSON response.
type Constructor struct {
	Id          string `json:"constructorId"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

//Driver contains driver information from the race schedule JSON response.
type Driver struct {
	Id          string `json:"driverId"`
	Url         string `json:"url"`
	Code        string `json:"code"`
	FirstName   string `json:"givenName"`
	LastName    string `json:"familyName"`
	DateOfBirth string `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
}

// DriverStanding contains the JSON response structure from a Get request
// to the driver standings URI.
type DriverStanding struct {
	Position     string        `json:"position"`
	Points       string        `json:"points"`
	Wins         string        `json:"wins"`
	DriverInfo   Driver        `json:"Driver"`
	Constructors []Constructor `json:"Constructors"`
}

// An entry in the standings list.
type standingsEntry struct {
	Season    string           `json:"season"`
	Round     string           `json:"round"`
	Standings []DriverStanding `json:"DriverStandings"`
}

// standingsTable contains the actual driver standings info for a season.
type standingsTable struct {
	Season        string           `json:"season"`
	StandingsList []standingsEntry `json:"StandingsLists"`
}

// mrData has header information for the driver standings.
type mrData struct {
	// todo: skipping a bunch of stuff.
	Table standingsTable `json:"StandingsTable"`
}

// standingsResponse is the top of the response tree for a Get request for
// the driver standings.
type standingsResponse struct {
	Data mrData `json:"MRData"`
}

// DriverStandings will return the current driver standings for the current season.
func DriverStandings() ([]DriverStanding, error) {
	const URL = "http://ergast.com/api/f1/current/driverStandings.json"
	var response standingsResponse
	err := request.Get(URL, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Table.StandingsList[0].Standings, nil
}
