package D_test

import (
	. "$ROOT_PATH$/D"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
