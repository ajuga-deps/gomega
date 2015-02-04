package matchers_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
	. "github.com/ajuga-deps/gomega/matchers"
)

var _ = Describe("BeTrue", func() {
	It("should handle true and false correctly", func() {
		Ω(true).Should(BeTrue())
		Ω(false).ShouldNot(BeTrue())
	})

	It("should only support booleans", func() {
		success, err := (&BeTrueMatcher{}).Match("foo")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())
	})
})
