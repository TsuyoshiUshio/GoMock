package f

import (
	"fmt"
)

type SecretValueGetter func(name string) string

func getKeyVault(name string) string {
	return fmt.Sprintf("Real: KeyVault Call %s", name)
}

func GetSecretValue(getter SecretValueGetter, name string) string {
	return getter(name)
}

// Usage Sample
func main() {
	fmt.Printf(GetSecretValue(getKeyVault, "SomeName"))
	fmt.Println()
}
