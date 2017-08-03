package s_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStruct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Struct Suite")
}
