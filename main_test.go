package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterPackages(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b"}

	diffResult := filterPackages(a, b)

	assert.Equal(t, []string{"c"}, diffResult)
}

func Test_filterPackagesWithPackageNames(t *testing.T) {
	a := []string{"github.com/bitrise-steplib/steps-go-list/gotool", "github.com/bitrise-steplib/steps-go-list"}
	b := []string{"github.com/bitrise-steplib/steps-go-list/gotool"}

	diffResult := filterPackages(a, b)

	assert.Equal(t, []string{"github.com/bitrise-steplib/steps-go-list"}, diffResult)
}
