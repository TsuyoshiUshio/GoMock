package i_test

import (
	"fmt"

	. "github.com/TsuyoshiUshio/GoMock/interface"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockKeyVaultClient struct {
}

func (c *MockKeyVaultClient) GetSecretValue(name string) string {
	fmt.Printf("**Fakeit")
	fmt.Println()
	return fmt.Sprintf("Fake: KeyVault Call %s", name)
}

var _ = Describe("Secret Value Client", func() {
	Context("When I request a secret", func() {

		It("returns secret value", func() {
			c := MockKeyVaultClient{}
			Expect("Fake: KeyVault Call SomeName", RetriveSecretValue(&c, "SomeValue"))
		})
	})
})
