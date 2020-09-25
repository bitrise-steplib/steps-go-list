package gotool

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/command"
)

// Commander ...
type Commander interface {
	ExecuteCommand(string, ...string) (string, error)
}

// CommandExecutor ...
type CommandExecutor struct{}

// ExecuteCommand ...
func (c CommandExecutor) ExecuteCommand(stringCommand string, args ...string) (string, error) {
	cmd := command.New(stringCommand, args...)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s failed: %s", cmd.PrintableCommandArgs(), out)
	}
	return out, nil
}

func parsePackages(out string) (list []string) {
	for _, l := range strings.Split(string(out), "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		list = append(list, l)
	}
	return list
}

// ListPackages ...
func ListPackages(commander Commander) ([]string, error) {
	executionResult, err := commander.ExecuteCommand("go", "list", "./...")
	if err != nil {
		return nil, err
	}

	return parsePackages(executionResult), nil
}
