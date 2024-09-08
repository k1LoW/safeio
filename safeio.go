package safeio

import (
	"io"
	"sync"
)

var _ io.Writer = (*Writer)(nil)

// Writer is a thread-safe io.Writer.
type Writer struct {
	wr io.Writer
	mu sync.Mutex
}

// NewWriter returns a new Writer.
func NewWriter(wr io.Writer) *Writer {
	return &Writer{
		wr: wr,
		mu: sync.Mutex{},
	}
}

// Write writes p to the underlying io.Writer.
func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.wr.Write(p)
}

// Switch switches the underlying io.Writer.
func (w *Writer) Switch(wr io.Writer) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wr = wr
}
