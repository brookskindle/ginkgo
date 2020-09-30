package suite_command_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestSuiteCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite Command Suite")
}
