package assertion_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"

	"testing"
)

func TestAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assertion Suite")
}
