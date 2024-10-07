package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer type from bytes package implements the Writer interface
	// a buffer is basically a []byte (slice of bytes)
	// basically, think of it like a String Builder

	// use a buffer for efficient I/O operations when reading or writing to data files,
	// network connections, or any stream where data arrives in parts
	// string concatenation is faster
	// also helps to create large byte sequences
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
