package buffer

import (
	"bytes"
	"sync"
)

// Pool is struct for buffer pool
type Pool struct {
	p *sync.Pool
}

// Buffer is main type for buffer
type Buffer struct {
	BS   *bytes.Buffer
	pool Pool
}

// NewPool Generate new buffer pool
func NewPool() Pool {
	return Pool{p: &sync.Pool{
		New: func() interface{} {
			var b bytes.Buffer
			return &Buffer{BS: &b}
		},
	}}
}

// Get new buffer from pool
func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.BS.Reset()
	buf.pool = p
	return buf
}

// put buffer back to pool
func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}

// Free will Release Buffer to pool
func (b *Buffer) Free() {
	b.pool.put(b)
}
