// A simple Circular Buffer (ring buffer) for bytes in Go.
package ring

import (
	"io"
)

// A Ring buffer containing a ring of bytes
type RingBuffer struct {
	b     []byte
	start int
	end   int
	size  int	
}

// Return true if the ring buffer is full, false otherwise.
func (rb *RingBuffer) IsFull() bool {
	return (rb.end+1)%len(rb.b) == rb.start
}

// Return true if the ring buffer is empty, false otherwise.
func (rb *RingBuffer) IsEmpty() bool {
	return rb.end == rb.start
}

func (rb *RingBuffer) GetSize() int {
	return rb.size
}

func (rb *RingBuffer) WriteByte(c byte) error {
	rb.b[rb.end] = c
	rb.end = (rb.end + 1) % len(rb.b)
	if rb.end == rb.start {
		rb.start = (rb.start + 1) % len(rb.b)
	}
	rb.size++
	return nil
}

func (rb *RingBuffer) ReadByte() (c byte, err error) {
	if rb.IsEmpty() {
		return rb.b[rb.start], io.EOF
	}
	c = rb.b[rb.start]
	rb.start = (rb.start + 1) % len(rb.b)
	rb.size--
	return c, nil
}

func (rb *RingBuffer) Read(p []byte) (n int, err error) {
    n = 0
    for i := 0; i < len(p); i++ {
        b, err := rb.ReadByte()
        if err != nil {
            return 0, err
        }
        p[i] = b
        n++
    }
    return
}

func (rb *RingBuffer) ReadNoChange(p []byte) (n int, err error) {
    start := rb.start
    size := rb.size
    n, err = rb.Read(p)
    rb.start = start
    rb.size = size
    return n, err
}

// Returns the content of the buffer without changing the next read byte.
func (rb *RingBuffer) ReadAhead() (p []byte, n int, err error) {
    start := rb.start
    p = make([]byte, len(rb.b))
    n, err = rb.Read(p)
    rb.start = start
    return p, n, err
}

func (rb *RingBuffer) Write(p []byte) (n int, err error) {
    for _, b := range p {
        rb.WriteByte(b)
    }
    return len(p), nil
}

// Create a new RingBuffer of the specified size.
func NewRingBuffer(size int) *RingBuffer {
	rb := new(RingBuffer)
	rb.b = make([]byte, size+1)
    rb.start = 0
    rb.end = 0
	return rb
}
