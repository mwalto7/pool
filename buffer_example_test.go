package pool_test

import (
	"fmt"
	"io"
	"strings"

	"github.com/mwalto7/pool/pool"
)

func ExampleGetBuffer() {
	// Get a bytes.Buffer from the pool.
	buf := pool.GetBuffer()
	defer buf.Close()

	// Do something with the buffer.
	io.Copy(buf, strings.NewReader("hello, world"))

	fmt.Println(buf.String())
	// Output:
	// hello, world
}
