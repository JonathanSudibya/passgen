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
	bs   *bytes.Buffer
	pool Pool
}

// NewPool Generate new buffer pool
func NewPool() Pool {
	return Pool{p: &sync.Pool{
		New: func() interface{} {
			var b bytes.Buffer
			return &Buffer{bs: &b}
		},
	}}
}

// Get new buffer from pool
func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.bs.Reset()
	buf.pool = p
	return buf
}

// put buffer back to pool
func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}

// AppendBytes will append array of byte to buffer
func (b *Buffer) AppendBytes(bs []byte) (int, error) {
	return b.bs.Write(bs)
}

// AppendByte will append a byte
func (b *Buffer) AppendByte(by byte) (int, error) {
	return 0, b.bs.WriteByte(by)
}

// AppendString will append string
func (b *Buffer) AppendString(s string) (int, error) {
	return b.bs.WriteString(s)
}

// Bytes return []byte
func (b *Buffer) Bytes() []byte {
	return b.bs.Bytes()
}

// Byte return byte[0]
func (b *Buffer) Byte() byte {
	return b.bs.Bytes()[0]
}

// String return string buffer
func (b *Buffer) String() string {
	return b.bs.String()
}

// Cap return buffer cap
func (b *Buffer) Cap() int {
	return b.bs.Cap()
}

// Len return buffer len
func (b *Buffer) Len() int {
	return b.bs.Len()
}

// Free will Release Buffer to pool
func (b *Buffer) Free() {
	b.pool.put(b)
}
