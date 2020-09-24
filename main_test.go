package main

import (
	"reflect"
	"testing"
)

func Test_filterPackages(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b"}

	diffResult := filterLines(a, b)

	filteredLines := FilteredLines{}
	filteredLines.Expected = []string{"c"}

	if !reflect.DeepEqual(filteredLines.Expected, diffResult.Expected) {
		t.Fatalf("expected result arrays should equal: expected: %s, recieved: %s", filteredLines.Expected, diffResult.Expected)
	}
}

func Test_ClearCommandResult(t *testing.T) {
	commandResult := []string{
		"go: downloading github.com/sirupsen/logrus v1.2.0",
		"go: downloading gopkg.in/alecthomas/kingpin.v2 v2.2.6",
		"go: downloading github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf",
		"go: downloading github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc",
		"go: downloading golang.org/x/crypto v0.0.0-20180904163835-0709b304e793",
		"go: downloading golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33",
		"github.com/skyrocknroll/go-mod-example",
	}

	clearedResult := filterLines(commandResult, commandExcludePatterns)

	filteredLines := FilteredLines{}
	filteredLines.Expected = []string{"github.com/skyrocknroll/go-mod-example"}

	if !reflect.DeepEqual(filteredLines.Expected, clearedResult.Expected) {
		t.Fatalf("expected result arrays should equal: expected: %s, recieved: %s", filteredLines.Expected, clearedResult.Expected)
	}
}
