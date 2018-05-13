package handler

import (
	"sync"
	"github.com/syyongx/llog/types"
)

type Buffer struct {
	Handler
	Processable

	BufferLimit     int
	FlushOnOverflow bool
	bufferSize      int
	h               types.IHandler
	buffer          chan *types.Record
	once            sync.Once
}

// New buffer handler
// bufferLimit: How many entries should be buffered at most, beyond that the oldest items are removed from the buffer.
// flushOnOverflow: If true, the buffer is flushed when the max size has been reached, by default oldest entries are discarded
func NewBuffer(handler types.IHandler, bufferLimit, level int, bubble, flushOnOverflow bool) *Buffer {
	buf := &Buffer{
		BufferLimit:     bufferLimit,
		FlushOnOverflow: flushOnOverflow,
		h:               handler,
		buffer:          make(chan *types.Record, bufferLimit),
		once:            sync.Once{},
	}
	buf.SetLevel(level)
	buf.SetBubble(bubble)
	return buf
}

func (buf *Buffer) Handle(record *types.Record) bool {
	if record.Level < buf.level {
		return false
	}
	if buf.BufferLimit > 0 && buf.bufferSize == buf.BufferLimit {
		if buf.FlushOnOverflow {
			buf.once.Do(func() {
				go buf.Flush()
			})
		} else {
			// If overflow remove the first record.
			<-buf.buffer
			//buf.bufferSize--
		}
	}
	if buf.processors != nil {
		buf.ProcessRecord(record)
	}
	//buf.buffer <- record
	select {
	case buf.buffer <- record:
	default:
		// channel is full
	}
	//buf.bufferSize++

	return false == buf.GetBubble()
}

func (buf *Buffer) HandleBatch(records chan *types.Record) {
	//buf.h.HandleBatch(records)
	for {
		record := <-records
		buf.h.Handle(record)
	}
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
	buf.buffer = nil
}

// close
func (buf *Buffer) Close() {
	//buf.h.Close()
	buf.Flush()
}
