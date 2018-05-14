package handler

import (
	"github.com/syyongx/llog/types"
	"sync"
)

type Buffer struct {
	Handler
	Processable

	FlushOnOverflow bool
	bufferLimit     int
	h               types.IHandler
	buffer          []*types.Record
	mu              sync.Mutex
}

// New buffer handler
// bufferLimit: How many entries should be buffered at most, beyond that the oldest items are removed from the buffer.
// flushOnOverflow: If true, the buffer is flushed when the max size has been reached, by default oldest entries are discarded
func NewBuffer(handler types.IHandler, bufferLimit, level int, bubble, flushOnOverflow bool) *Buffer {
	b := &Buffer{
		FlushOnOverflow: flushOnOverflow,
		bufferLimit:     bufferLimit,
		h:               handler,
		buffer:          make([]*types.Record, 0, bufferLimit),
	}
	b.SetLevel(level)
	b.SetBubble(bubble)
	return b
}

func (b *Buffer) Handle(record *types.Record) bool {
	if record.Level < b.level {
		return false
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.bufferLimit > 0 && len(b.buffer) == b.bufferLimit {
		if b.FlushOnOverflow {
			b.Flush()
		} else {
			// If overflow remove the first record.
			b.buffer = b.buffer[1:]
		}
	}
	if b.processors != nil {
		b.ProcessRecord(record)
	}
	b.buffer = append(b.buffer, record)

	return false == b.GetBubble()
}

func (b *Buffer) HandleBatch(records []*types.Record) {
	b.h.HandleBatch(records)
}

func (b *Buffer) Flush() {
	if len(b.buffer) == 0 {
		return
	}
	b.HandleBatch(b.buffer)
	b.Clear()
}

// Clears the buffer without flushing any messages down to the wrapped handler.
func (b *Buffer) Clear() {
	b.buffer = b.buffer[:0]
}

// close
func (b *Buffer) Close() {
	//b.h.Close()
	b.Flush()
}
