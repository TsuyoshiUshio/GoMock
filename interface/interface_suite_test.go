package i_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInterface(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interface Suite")
}
