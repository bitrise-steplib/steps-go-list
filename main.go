package main

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-go-list/gotool"
	"github.com/bitrise-tools/go-steputils/tools"
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

func filterPackages(original, exclude []string) []string {
	m := make(map[string]bool)
	for _, s := range original {
		m[s] = true
	}
	for _, s := range exclude {
		m[s] = false
	}

	var result []string
	for k, v := range m {
		if v {
			log.Donef("✓ %s", k)
			result = append(result, k)
		} else {
			log.Printf("- %s", k)
		}
	}

	return result
}
