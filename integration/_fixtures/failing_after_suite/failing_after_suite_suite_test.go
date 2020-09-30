package failing_before_suite_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestFailingAfterSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FailingAfterSuite Suite")
}

var _ = BeforeSuite(func() {
	println("BEFORE SUITE")
})

var _ = AfterSuite(func() {
	println("AFTER SUITE")
	panic("BAM!")
})
