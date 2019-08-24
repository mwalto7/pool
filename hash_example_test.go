package pool_test

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/mwalto7/pool/pool"
)

func ExampleGetHash_sha256() {
	// Get a SHA256 hash function from the pool.
	h := pool.GetHash(crypto.SHA256)
	defer h.Close() // Dont forget to call Close! This puts the hash function back into the pool.

	// Hash something.
	io.Copy(h, strings.NewReader("hello, world"))

	// Get the hash sum.
	sum := h.Sum(nil)
	fmt.Println("SHA256:", hex.EncodeToString(sum))
	// Output:
	// SHA256: 09ca7e4eaa6e8ae9c7d261167129184883644d07dfba7cbfbc4c8a2e08360d5b
}

func ExampleGetHash_multiple() {
	// Get a SHA256 hash function from the pool.
	sha256 := pool.GetHash(crypto.SHA256)
	defer sha256.Close()

	// Get an MD5 hash function from the pool.
	md5 := pool.GetHash(crypto.MD5)
	defer md5.Close()

	// Hash something with the hash functions.
	io.Copy(io.MultiWriter(sha256, md5), strings.NewReader("hello, world"))

	// Get the hash sums.
	fmt.Println("SHA256:", hex.EncodeToString(sha256.Sum(nil)))
	fmt.Println("MD5:", hex.EncodeToString(md5.Sum(nil)))
	// Output:
	// SHA256: 09ca7e4eaa6e8ae9c7d261167129184883644d07dfba7cbfbc4c8a2e08360d5b
	// MD5: e4d7f1b4ed2e42d15898f4b27b019da4
}
