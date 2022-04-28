package test

import (
	"fmt"
	"testing"

	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
)

func TestConstructorStandings(t *testing.T) {
	standings, err := ergast.ConstructorStandings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Constructor Standings:\n%v\n", standings)
}
