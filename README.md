Three Mock storategy using Go
===
As an extreme programmer, I'd like to learn BDD/TDD with go. Mock is an important technique for the methods. I'd like to write three method for Mocking with Go. You might not need to mock via mock framework. 

# 1. Mock by Function

You can mock by function. Just inject a function which has the same parameter type and return value type. 

_function.go_

```
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
```

_function_test.go_

```
package f_test

import (
	"fmt"

	. "github.com/TsuyoshiUshio/GoMock/cmd"
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
```

As you can see, `mock_get_keyVault` has been passed to the `GetSecretValue`.

Result is as expected.

```
$ ginkgo cmd
Running Suite: Cmd Suite
========================
Random Seed: 1501736262
Will run 1 of 1 specs

**Fakeit
•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

Ginkgo ran 1 suite in 653.771446ms
Test Suite Passed
```

# 2. Mock by struct field

Almost the same, however, Go can treat struct like Class oriented language. 

_struct.go_

```
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

```

_struct_test.go_

```
package s_test

import (
	"fmt"

	. "github.com/TsuyoshiUshio/GoMock/struct"
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
			c := NewKeyVaultClient(mock_getKeyVault)
			Expect("Fake: KeyVault Call SomeName", c.GetSecretValue("SomeName"))
		})
	})
})
```

Result

```
$ go test -v struct/*.go
=== RUN   TestStruct
Running Suite: Struct Suite
===========================
Random Seed: 1501737616
Will run 1 of 1 specs

**Fakeit
•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped --- PASS: TestStruct (0.00s)
PASS
ok  	command-line-arguments	0.011s
```
As a production code relatively complex. At that time, the second strategy works. You can pass the function as a field of struct.

# 3. Mock by interface

You can also use interface. 

_interface.go_

```
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
```

_interface_test.go_

```
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
```

This is similar to the Class/Inteface Mocking style. If you create an interface, the type should implement the target function. 

Result

```
$ ginkgo interface/
Running Suite: Interface Suite
==============================
Random Seed: 1501738567
Will run 1 of 1 specs

**Fakeit
•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

Ginkgo ran 1 suite in 682.091566ms
Test Suite Passed
```

# Conclusion

I learnt how to mock the go programming. Now I'm ready to Test Driven Development with Go!


