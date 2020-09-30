package test_description_test

import (
	"fmt"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("TestDescription", func() {
	It("should pass", func() {
		Ω(true).Should(BeTrue())
	})

	It("should fail", func() {
		Ω(true).Should(BeFalse())
	})

	AfterEach(func() {
		description := CurrentGinkgoTestDescription()
		fmt.Printf("%s:%t\n", description.FullTestText, description.Failed)
	})
})
