# Ginkgo testing helpers
Ginkgo testing helpers to make infrastructure tests simpler and more comprehensible.

## Usage

Import this module into your Ginkgo/Gomega based testing suites.

```go
import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/sueddeutsche/plat-testing-helpers"
)

var _ = Describe("my test context", func() {
  longString := "hello big and nice world"
  It("should do a test case", func() {
		ExpectSubstrings(longString, "hello", "nice", "world") // passes
		ExpectMissingSubstrings(longString, "bad people") // passes
	})
})
```

```shell
go mod tidy && go test
```