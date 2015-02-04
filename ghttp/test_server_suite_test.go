package ghttp_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"

	"testing"
)

func TestGHTTP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GHTTP Suite")
}
