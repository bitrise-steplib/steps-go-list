package gotool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ListPackages2(t *testing.T) {
	t.Run("Test list packages", func(t *testing.T) {
		expectedResult := []string{
			"github.com/bitrise-steplib/steps-go-list",
			"github.com/bitrise-steplib/steps-go-list/gotool",
		}

		mockCommander := GivenMockCommander()
		mockCommander.
			On("ExecuteCommand", mock.Anything, mock.Anything).
			Return(
				"github.com/bitrise-steplib/steps-go-list\ngithub.com/bitrise-steplib/steps-go-list/gotool", nil)

		packages, err := ListPackages(mockCommander)

		require.NoError(t, err)
		assert.Equal(t, expectedResult, packages)
		mockCommander.AssertExpectations(t)
	})

	t.Run("Test list packages with empty line", func(t *testing.T) {
		expectedResult := []string{
			"github.com/bitrise-steplib/steps-go-list",
			"github.com/bitrise-steplib/steps-go-list/gotool",
		}

		mockCommander := GivenMockCommander()
		mockCommander.
			On("ExecuteCommand", mock.Anything, mock.Anything).
			Return(
				"github.com/bitrise-steplib/steps-go-list\ngithub.com/bitrise-steplib/steps-go-list/gotool\n", nil)

		packages, err := ListPackages(mockCommander)

		require.NoError(t, err)
		assert.Equal(t, expectedResult, packages)
		mockCommander.AssertExpectations(t)
	})

	t.Run("Package list error", func(t *testing.T) {
		resultError := errors.New("Listing error")
		mockCommander := GivenMockCommander()
		mockCommander.
			On("ExecuteCommand", mock.Anything, mock.Anything).
			Return(
				"", resultError)

		packages, err := ListPackages(mockCommander)

		require.EqualError(t, err, resultError.Error())
		assert.Equal(t, []string(nil), packages)
		mockCommander.AssertExpectations(t)
	})
}
