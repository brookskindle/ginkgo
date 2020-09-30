package C_test

import (
	. "$ROOT_PATH$/C"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
