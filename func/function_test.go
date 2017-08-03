package f_test

import (
	"fmt"

	. "github.com/TsuyoshiUshio/GoMock/func"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func mock_getKeyVault(name string) string {
	fmt.Printf("**Fakeit")
	fmt.Println()
	return fmt.Sprintf("Fake: KeyVault Call %s", name)
}

var _ = Describe("Secret Value Client", func() {
	Context("When I request a secret", func() {

		It("returns secret value", func() {
			Expect("Fake: KeyVault Call SomeName", GetSecretValue(mock_getKeyVault, "SomeName"))
		})
	})
})
