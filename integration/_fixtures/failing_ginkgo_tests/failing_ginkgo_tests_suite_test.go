package failing_ginkgo_tests_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestFailing_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failing_ginkgo_tests Suite")
}
