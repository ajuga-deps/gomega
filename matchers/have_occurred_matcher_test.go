package matchers_test

import (
	"errors"
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
	. "github.com/ajuga-deps/gomega/matchers"
)

var _ = Describe("HaveOccurred", func() {
	It("should succeed if matching an error", func() {
		Ω(errors.New("Foo")).Should(HaveOccurred())
	})

	It("should not succeed with nil", func() {
		Ω(nil).ShouldNot(HaveOccurred())
	})

	It("should only support errors and nil", func() {
		success, err := (&HaveOccurredMatcher{}).Match("foo")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())

		success, err = (&HaveOccurredMatcher{}).Match("")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())
	})
})
