package handler

import (
	"github.com/syyongx/llog/types"
)

type Buffer struct {
	Handler
	Processable

	h               types.IHandler
	bufferSize      int
	bufferLimit     int
	flushOnOverflow bool
	buffer          []*types.Record
}

// New buffer handler
// bufferLimit: How many entries should be buffered at most, beyond that the oldest items are removed from the buffer.
// flushOnOverflow: If true, the buffer is flushed when the max size has been reached, by default oldest entries are discarded
func NewBuffer(handler types.IHandler, bufferLimit, level int, bubble, flushOnOverflow bool) *Buffer {
	buf := &Buffer{
		h:               handler,
		bufferLimit:     bufferLimit,
		flushOnOverflow: flushOnOverflow,
		buffer:          make([]*types.Record, 0, bufferLimit),
	}
	buf.SetLevel(level)
	buf.SetBubble(bubble)
	return buf
}

func (buf *Buffer) Handle(record *types.Record) bool {
	if record.Level < buf.level {
		return false
	}
	if buf.bufferLimit > 0 && buf.bufferSize == buf.bufferLimit {
		if buf.flushOnOverflow {
			buf.Flush()
		} else {
			// If overflow remove the first record.
			buf.buffer = buf.buffer[1:]
			buf.bufferSize--
		}
	}
	if buf.processors != nil {
		buf.ProcessRecord(record)
	}
	buf.buffer = append(buf.buffer, record)
	buf.bufferSize++

	return false == buf.GetBubble()
}

func (buf *Buffer) HandleBatch(records []*types.Record) {
	buf.h.HandleBatch(records)
}

func (buf *Buffer) Flush() {
	if buf.bufferSize == 0 {
		return
	}
	buf.HandleBatch(buf.buffer)
}

// Clears the buffer without flushing any messages down to the wrapped handler.
func (buf *Buffer) Clear() {
	buf.bufferSize = 0
	buf.buffer = buf.buffer[:0]
}

// close
func (buf *Buffer) Close() {
	//buf.h.Close()
	buf.Flush()
}
