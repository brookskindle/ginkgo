package debug_parallel_fixture_test

import (
	"testing"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

func TestDebugParallelFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DebugParallelFixture Suite")
}
