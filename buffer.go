package pool

import (
	"bytes"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return &BuffCloser{new(bytes.Buffer)}
	},
}

// GetBuffer returns a pooled bytes.Buffer and a reset function.
func GetBuffer() *BuffCloser {
	return buffers.Get().(*BuffCloser)
}

// BuffCloser represents a pooled bytes.Buffer.
type BuffCloser struct {
	*bytes.Buffer
}

// Close resets the buffer and puts the buffer back into the pool.
func (b *BuffCloser) Close() error {
	if b != nil && b.Buffer != nil {
		b.Buffer.Reset()
		buffers.Put(b)
	}
	return nil
}
