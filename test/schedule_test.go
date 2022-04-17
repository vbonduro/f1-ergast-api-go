package test

import (
	"fmt"
	"testing"

	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
)

func TestCurrentSchedule(t *testing.T) {
	races, err := ergast.RaceSchedule()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Race Schedule:\n%v\n", races)
}

func TestNextRace(t *testing.T) {
	race, err := ergast.NextRace()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Next Race:\n%v\n", race)
}
