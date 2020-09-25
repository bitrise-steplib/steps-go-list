package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterPackages(t *testing.T) {
	original := []string{"same", "same_too", "Expected"}
	excludes := []string{"same", "same_too"}

	diffResult := filterLines(original, excludes)

	filteredLines := FilteredLines{}
	filteredLines.Expected = []string{"Expected"}
	filteredLines.Removed = []string{"same", "same_too"}

	assert.Equal(t, filteredLines.Expected, diffResult.Expected)
	assert.Equal(t, filteredLines.Removed, diffResult.Removed)
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

	assert.Equal(t, filteredLines.Expected, clearedResult.Expected)
}
