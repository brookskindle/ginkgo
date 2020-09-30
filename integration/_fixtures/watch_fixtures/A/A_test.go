package A_test

import (
	. "$ROOT_PATH$/A"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
