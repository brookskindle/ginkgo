package D_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}
