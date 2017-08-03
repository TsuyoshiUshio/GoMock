package s

import (
	"fmt"
)

type SecretValueGetter func(name string) string

type KeyVaultClient struct {
	get_value SecretValueGetter
}

func NewKeyVaultClient(getter SecretValueGetter) *KeyVaultClient {
	return &KeyVaultClient{get_value: getter}
}

func (c *KeyVaultClient) GetSecretValue(name string) string {
	return c.get_value(name)
}

func getKeyVault(name string) string {
	return fmt.Sprintf("Real: KeyVault Call %s", name)
}

func main() {
	c := NewKeyVaultClient(getKeyVault)
	fmt.Printf(c.GetSecretValue("SomeName"))
	fmt.Println()
}
