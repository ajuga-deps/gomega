package testingtsupport_test

import (
	. "github.com/ajuga-deps/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
