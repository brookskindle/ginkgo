package failing_ginkgo_tests_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/ginkgo/integration/_fixtures/failing_ginkgo_tests"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("FailingGinkgoTests", func() {
	It("should fail", func() {
		Ω(AlwaysFalse()).Should(BeTrue())
	})

	It("should pass", func() {
		Ω(AlwaysFalse()).Should(BeFalse())
	})
})
