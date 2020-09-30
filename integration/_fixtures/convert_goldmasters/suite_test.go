package tmp_test

import (
	"testing"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

func TestConvertFixtures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertFixtures Suite")
}
