package nested_test

import (
	"testing"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}
