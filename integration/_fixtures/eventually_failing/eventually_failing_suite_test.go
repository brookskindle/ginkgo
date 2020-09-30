package eventually_failing_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestEventuallyFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventuallyFailing Suite")
}
