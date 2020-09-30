package table_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Table Suite")
}
