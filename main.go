package main

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-go-list/gotool"
	"github.com/bitrise-tools/go-steputils/tools"
	"github.com/ryanuber/go-glob"
)

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

	commandExecutor := gotool.CommandExecutor{}
	packages, err := gotool.ListPackages(commandExecutor)
	if err != nil {
		failf("Failed to list packages: %s", err)
	}

	log.Infof("\nList of packages:")
	filteredPackages := filterPackages(packages, excludes)

	if err := tools.ExportEnvironmentWithEnvman("BITRISE_GO_PACKAGES", strings.Join(filteredPackages, "\n")); err != nil {
		failf("Failed to export packages, error: %s", err)
	}
}

func filterPackages(original, excludes []string) []string {
	var result []string
	for _, p := range original {
		if !matching(p, excludes) {
			log.Donef("✓ %s", p)
			result = append(result, p)
		} else {
			log.Printf("- %s", p)
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
