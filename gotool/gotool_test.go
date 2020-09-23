package gotool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCommander struct {
	ExecuteCommandFn func(command string, args ...string) (string, error)
}

func (c TestCommander) ExecuteCommand(command string, args ...string) (string, error) {
	if c.ExecuteCommandFn == nil {
		panic("You have to override TestCommander.ExecuteCommand function in tests")
	}
	return c.ExecuteCommandFn(command, args...)
}

func Test_ListPackages(t *testing.T) {
	testCases := []struct {
		desc               string
		executionResultRaw string
		expectedError      error
		expectedResult     []string
		commander          Commander
	}{
		{
			desc:          "Test list packages",
			expectedError: nil,
			expectedResult: []string{
				"github.com/bitrise-steplib/steps-go-list",
				"github.com/bitrise-steplib/steps-go-list/gotool",
			},
			commander: TestCommander{
				ExecuteCommandFn: func(command string, args ...string) (string, error) {
					return "github.com/bitrise-steplib/steps-go-list\ngithub.com/bitrise-steplib/steps-go-list/gotool", nil
				},
			},
		},
		{
			desc:          "Test list packages with empty line",
			expectedError: nil,
			expectedResult: []string{
				"github.com/bitrise-steplib/steps-go-list",
				"github.com/bitrise-steplib/steps-go-list/gotool",
			},
			commander: TestCommander{
				ExecuteCommandFn: func(command string, args ...string) (string, error) {
					return "github.com/bitrise-steplib/steps-go-list\ngithub.com/bitrise-steplib/steps-go-list/gotool\n", nil
				},
			},
		},
		{
			desc:           "Package list error",
			expectedError:  errors.New("Listing error"),
			expectedResult: nil,
			commander: TestCommander{
				ExecuteCommandFn: func(command string, args ...string) (string, error) {
					return "", errors.New("Listing error")
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			packages, err := ListPackages(tC.commander)

			assert.Equal(t, tC.expectedError, err)
			assert.Equal(t, tC.expectedResult, packages)
		})
	}
}
