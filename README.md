# pool
[![GoDoc](https://godoc.org/github.com/mwalto7/pool?status.svg)](https://godoc.org/github.com/mwalto7/pool)
[![Go Report Card](https://goreportcard.com/badge/github.com/mwalto7/pool)](https://goreportcard.com/report/github.com/mwalto7/pool)

## Getting Started

```
go get -u github.com/mwalto7/pool
```

## Example

```go
package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/mwalto7/pool"
)

func main() {
    // Get a bytes.Buffer from the pool.
    buf := pool.GetBuffer()
    defer buf.Close() // Put the bytes.Buffer back into the pool!

    // Get a SHA256 hash function from the pool.
    h := pool.GetHash(crypto.SHA256)
    defer h.Close() // Put the hash function back into the pool!

    // Do something with the buffer and hash function.
    io.Copy(io.MultiWriter(buf, h), strings.NewReader("hello, world"))

    // Print the results.
    fmt.Printf("SHA256 %q: %s", buf.String(), hex.EncodeToString(h.Sum(nil)))
}
```

```
Output:
SHA256 "hello, world": 09ca7e4eaa6e8ae9c7d261167129184883644d07dfba7cbfbc4c8a2e08360d5b
```
