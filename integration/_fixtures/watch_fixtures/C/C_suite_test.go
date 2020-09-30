package C_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
