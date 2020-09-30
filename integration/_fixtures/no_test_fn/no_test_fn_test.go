package no_test_fn_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/ginkgo/integration/_fixtures/no_test_fn"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("NoTestFn", func() {
	It("should proxy strings", func() {
		Ω(StringIdentity("foo")).Should(Equal("foo"))
	})
})
