#chilru

chilru is lru cache, written in go

# Install

```shell
$ go get github.com/langwan/chilru
```

# Example

```go
package main

import (
	"fmt"
	"github.com/langwan/chilru"
)

func main() {
	lru := chilru.New[string, int](100)
	lru.Add("a", 1)
	lru.Set("b", 2)
	lru.Range(func(k, v any) bool {
		fmt.Println("k", k)
		fmt.Println("v", v)
		return true
	})
}

```