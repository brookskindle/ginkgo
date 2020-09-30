package B_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}
