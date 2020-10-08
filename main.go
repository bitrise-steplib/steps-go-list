package main

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-go-list/gotool"
	"github.com/bitrise-tools/go-steputils/tools"
	"github.com/ryanuber/go-glob"
)

// FilteredLines ...
type FilteredLines struct {
	Removed  []string
	Expected []string
}

var commandExcludePatterns = []string{
	"*go: downloading*",
}

func failf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {
	exclude := os.Getenv("exclude")

	log.Infof("Configs:")
	log.Printf("- exclude: %s", exclude)

	if exclude == "" {
		failf("Required input not defined: exclude")
	}

	excludes := strings.Split(exclude, "\n")

	commandResult, err := gotool.ListPackages(gotool.CommandExecutor{})
	if err != nil {
		failf("Failed to list packages: %s", err)
	}

	filteredOutput := filterLines(commandResult, commandExcludePatterns)
	for _, l := range filteredOutput.Removed {
		log.Printf("%s", l)
	}

	log.Infof("\nList of packages:")
	filteredPackages := filterLines(filteredOutput.Expected, excludes)
	for _, l := range filteredPackages.Removed {
		log.Printf("- %s", l)
	}
	for _, l := range filteredPackages.Expected {
		log.Donef("✓ %s", l)
	}

	if err := tools.ExportEnvironmentWithEnvman("BITRISE_GO_PACKAGES", strings.Join(filteredPackages.Expected, "\n")); err != nil {
		failf("Failed to export packages, error: %s", err)
	}
}

func filterLines(original, excludes []string) FilteredLines {
	var result FilteredLines
	for _, p := range original {
		if matching(p, excludes) {
			result.Removed = append(result.Removed, p)
		} else {
			result.Expected = append(result.Expected, p)
		}
	}

	return result
}

func matching(str string, matches []string) bool {
	for _, e := range matches {
		if glob.Glob(e, str) {
			return true
		}
	}

	return false
}
