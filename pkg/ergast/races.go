package ergast

import (
	"strconv"
	"time"

	"github.com/vbonduro/f1-ergast-api-go/internal/request"
	"github.com/vbonduro/f1-ergast-api-go/internal/response"
)

// GeoCoordinate describes the location where the circuit is.
type GeoCoordinate struct {
	Latitude  float64
	Longitude float64
}

// Circuit contains information about an F1 Circuit.
type Circuit struct {
	Name     string
	Location GeoCoordinate
	Country  string
	Locality string
}

// Race has all pertinent infromation about an F1 Race.
type Race struct {
	RoundNumber         int
	Name                string
	CircuitInfo         Circuit
	FirstPracticeStart  time.Time
	SprintRaceStart     time.Time
	SecondPracticeStart time.Time
	ThirdPracticeStart  time.Time
	QualifyingStart     time.Time
	RaceStart           time.Time
}

// RaceSchedule will return the race schedule for the current year.
func RaceSchedule() ([]Race, error) {
	const URL = "http://ergast.com/api/f1/current.json"
	return getRaces(URL)
}

func NextRace() (*Race, error) {
	const URL = "http://ergast.com/api/f1/current/next.json"
	races, err := getRaces(URL)
	if err != nil {
		return nil, err
	}
	return &races[0], nil
}

func getRaces(uri string) ([]Race, error) {
	var scheduleResponse response.ScheduleResponse
	err := request.Get(uri, &scheduleResponse)
	if err != nil {
		return nil, err
	}

	var races []Race
	for _, schedule := range scheduleResponse.Header.ScheduleTable.Races {
		race, err := makeRace(schedule)
		if err != nil {
			return nil, err
		}
		races = append(races, *race)
	}

	return races, nil
}

func makeRace(response response.RaceSchedule) (*Race, error) {
	round, err := strconv.Atoi(response.Round)
	if err != nil {
		return nil, err
	}

	var race Race
	race.RoundNumber = round
	race.Name = response.Name

	circuit, err := makeCircuit(response.Circuit)
	if err != nil {
		return nil, err
	}
	race.CircuitInfo = *circuit

	err = populateTimestamps(&race, &response)
	if err != nil {
		return nil, err
	}

	return &race, nil
}

func makeCircuit(response response.RaceCircuit) (*Circuit, error) {
	var circuit Circuit
	circuit.Country = response.LocationInfo.Country
	circuit.Locality = response.LocationInfo.Locality
	circuit.Name = response.CircutName

	longitude, err := strconv.ParseFloat(response.LocationInfo.Longitude, 64)
	if err != nil {
		return nil, err
	}
	latitude, err := strconv.ParseFloat(response.LocationInfo.Latitude, 64)
	if err != nil {
		return nil, err
	}

	circuit.Location.Longitude = longitude
	circuit.Location.Latitude = latitude

	return &circuit, nil
}

func populateTimestamps(race *Race, response *response.RaceSchedule) error {
	raceTime, err := makeTime(response.Date, response.Time)
	if err != nil {
		return err
	}

	practice1Time, err := makeTime(response.FirstPractice.Date, response.FirstPractice.Time)
	if err != nil {
		return err
	}

	sprintTime, err := makeTime(response.Sprint.Date, response.Sprint.Time)
	if err != nil {
		return err
	}

	practice2Time, err := makeTime(response.SecondPractice.Date, response.SecondPractice.Time)
	if err != nil {
		return err
	}

	practice3Time, err := makeTime(response.ThirdPractice.Date, response.ThirdPractice.Time)
	if err != nil {
		return err
	}

	qualifyingTime, err := makeTime(response.Qualifying.Date, response.Qualifying.Time)
	if err != nil {
		return err
	}

	race.RaceStart = raceTime
	race.FirstPracticeStart = practice1Time
	race.SprintRaceStart = sprintTime
	race.SecondPracticeStart = practice2Time
	race.ThirdPracticeStart = practice3Time
	race.QualifyingStart = qualifyingTime

	return nil
}

func makeTime(date string, timeofday string) (time.Time, error) {
	if len(date) == 0 || len(timeofday) == 0 {
		return time.Time{}, nil
	}
	rfc3339Timestamp := date + "T" + timeofday
	return time.Parse(time.RFC3339, rfc3339Timestamp)
}
