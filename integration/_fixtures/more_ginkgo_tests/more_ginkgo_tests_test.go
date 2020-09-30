package more_ginkgo_tests_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/ginkgo/integration/_fixtures/more_ginkgo_tests"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
