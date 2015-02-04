package asyncassertion_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"

	"testing"
)

func TestAsyncAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AsyncAssertion Suite")
}
