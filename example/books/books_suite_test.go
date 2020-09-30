package books_test

import (
	"testing"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
