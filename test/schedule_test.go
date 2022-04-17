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
	fmt.Printf("%v\n", races)
}
