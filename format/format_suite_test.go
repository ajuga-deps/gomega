package format_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
