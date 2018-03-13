package buffer

import (
	"sync"
)

// Pool is struct for buffer pool
type Pool struct {
	p *sync.Pool
}

// Buffer is main type for buffer
type Buffer struct {
	BS   []byte
	len  int
	pool Pool
}

// NewPool Generate new buffer pool
func NewPool() Pool {
	return Pool{p: &sync.Pool{
		New: func() interface{} {
			return &Buffer{BS: make([]byte, 0, 1024)}
		},
	}}
}

// Get new buffer from pool
func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.Reset()
	buf.pool = p
	return buf
}

// put buffer back to pool
func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}

// Reset bytes buffer
func (b *Buffer) Reset() {
	b.BS = b.BS[:0]
	b.len = 0
}

func (b *Buffer) Grow(n int) {
	b.len += n
	b.BS = b.BS[:b.len]
}

// Write bytes to buffer
func (b *Buffer) Write(bs []byte) (int, error) {
	m := b.len
	n := m + len(bs)
	b.Grow(len(bs))
	return copy(b.BS[m:n], bs), nil
}

func (b *Buffer) Len() int {
	return b.len
}

// WriteByte will write a Byte
func (b *Buffer) WriteByte(v byte) error {
	b.BS = append(b.BS, v)
	return nil
}

// WriteString will string string as []byte
func (b *Buffer) WriteString(s string) error {
	b.Write([]byte(s))
	return nil
}

// Bytes return current bytes
func (b *Buffer) Bytes() []byte {
	return b.BS
}

func (b *Buffer) String() string {
	return string(b.BS)
}

// Free will Release Buffer to pool
func (b *Buffer) Free() {
	b.pool.put(b)
}
