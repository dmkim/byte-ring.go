package ring

import (
	"testing"
)

func TestRing(t *testing.T) {
	ring := NewRingBuffer(10)
	if !ring.IsEmpty() {
		t.Fatal("Buffer not empty.")
	}
	ring.WriteByte(1)
	if ring.IsEmpty() {
		t.Fatal("Buffer empty.")
	}
	ring.WriteByte(2)
	ring.WriteByte(3)
	ring.WriteByte(4)
	ring.WriteByte(5)
	ring.WriteByte(6)
	ring.WriteByte(7)
	ring.WriteByte(8)
	ring.WriteByte(9)
	if ring.IsEmpty() {
		t.Fatal("Buffer empty.")
	}
	if ring.IsFull() {
		t.Fatal("Buffer full.")
	}
	ring.WriteByte(10)
	if ring.IsEmpty() {
		t.Fatal("Buffer empty.")
	}
	if !ring.IsFull() {
		t.Fatal("Buffer should be full but is not.")
	}
	ring.WriteByte(11)
	ring.WriteByte(12)
	b, err := ring.ReadByte()
	if err != nil {
		t.Fatalf("Error %s while reading", err)
	}
	if ring.IsFull() {
		t.Fatal("Buffer should not be full.")
	}
	if b != 3 {
		t.Fatalf("b (%d) != 3", b)
	}
}
