package tmp

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestTmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmp Suite")
}
