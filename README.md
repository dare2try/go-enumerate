# go-enumerate
Provides enumerating capabilities for golang data structures

## Installation

```bash
$ go get github.com/dare2try/go-enumerate
```

## Examples

```go
import (
    "fmt"
    "github.com/dare2try/go-enumerate"
)

func main() {
    // iterate over a slice
    a := []int{1, 2, 3}
    iterator := enumerate.Slice(a)
    for item, ok := iterator.Next(); ok; item, ok = iterator.Next() {
        fmt.Printf("Item: %s\n", item)
    }

    // iterate over a map
    b := []int{"A":1, "B":2, "C":3}
    iterator := enumerate.Map(b)
    for key, value, ok := iterator.Next(); ok; key, value, ok = iterator.Next() {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }
}
```
