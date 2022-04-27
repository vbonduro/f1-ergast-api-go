package test

import (
	"fmt"
	"testing"

	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
)

func TestDriverStandings(t *testing.T) {
	standings, err := ergast.DriverStandings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Driver Standings:\n%v\n", standings)
}
