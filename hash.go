package pool

import (
	"crypto"
	"hash"
	"sync"

	// Imported to make sure each package runs crypto.RegisterHash.
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	// Imported to make sure each package runs crypto.RegisterHash.
	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

func init() {
	registerHashFunc(crypto.MD4)
	registerHashFunc(crypto.MD5)
	registerHashFunc(crypto.SHA1)
	registerHashFunc(crypto.SHA224)
	registerHashFunc(crypto.SHA256)
	registerHashFunc(crypto.SHA384)
	registerHashFunc(crypto.SHA512)
	registerHashFunc(crypto.MD5SHA1)
	registerHashFunc(crypto.RIPEMD160)
	registerHashFunc(crypto.SHA3_224)
	registerHashFunc(crypto.SHA3_256)
	registerHashFunc(crypto.SHA3_384)
	registerHashFunc(crypto.SHA3_512)
	registerHashFunc(crypto.SHA512_224)
	registerHashFunc(crypto.SHA512_256)
	registerHashFunc(crypto.BLAKE2s_256)
	registerHashFunc(crypto.BLAKE2b_256)
	registerHashFunc(crypto.BLAKE2b_384)
	registerHashFunc(crypto.BLAKE2b_512)
}

var registeredHashFuncs = make(map[crypto.Hash]*hashPool)

// registerHashFunc registers the given hash function.
//
// This function should only be called during initialization (i.e. in an init()
// function). If more than one Hasher is registered with the same name, the last
// one to be registered will take effect.
func registerHashFunc(h crypto.Hash) {
	p := &hashPool{}
	p.hashFuncs.New = func() interface{} {
		return &hashFunc{Hash: h.New(), pool: &p.hashFuncs}
	}
	registeredHashFuncs[h] = p
}

// GetHash returns a pooled hash function.
func GetHash(h crypto.Hash) HashCloser {
	return registeredHashFuncs[h].getHash()
}

type hashPool struct {
	hashFuncs sync.Pool
}

func (p *hashPool) getHash() HashCloser {
	return p.hashFuncs.Get().(*hashFunc)
}

// HashCloser represents a pooled hash.Hash.
type HashCloser interface {
	hash.Hash
	// Close resets the hash to its initial state and puts it back into the pool.
	Close() error
}

type hashFunc struct {
	hash.Hash
	pool *sync.Pool
}

func (h *hashFunc) Close() error {
	if h != nil && h.Hash != nil && h.pool != nil {
		h.Hash.Reset()
		h.pool.Put(h)
	}
	return nil
}
