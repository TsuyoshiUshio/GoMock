package i

import "fmt"

type IKeyVaultClient interface {
	GetSecretValue(string) string
}

type KeyVaultClient struct {
}

func (c *KeyVaultClient) GetSecretValue(name string) string {
	return fmt.Sprintf("Real: KeyVault Call %s", name)
}

func RetriveSecretValue(c IKeyVaultClient, name string) string {
	return c.GetSecretValue(name)
}

func main() {
	c := KeyVaultClient{}
	fmt.Printf(RetriveSecretValue(&c, "SomeValue"))
	fmt.Println()
}
