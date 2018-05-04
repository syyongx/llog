package handler

import "github.com/syyongx/llog/types"

type Buffer struct {
	Handler
	Processable

	bufferSize      int
	bufferLimit     int
	flushOnOverflow bool
	buffer          []types.Record
}

func NewBuff() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Handle(record *types.Record) bool {
	if record.Level < b.level {
		return false
	}
	if b.bufferLimit > 0 && b.bufferSize == b.bufferLimit {
		if b.flushOnOverflow {
			b.Flush()
		} else {
			b.bufferSize--
		}
	}
	if b.processors != nil {
		record = b.ProcessRecord(record)
	}

	return false == b.bubble
}

func (b *Buffer) Flush() {
	if b.bufferSize == 0 {
		return
	}
	//b.HandleBatch(b.buffer)
}

// Clears the buffer without flushing any messages down to the wrapped handler.
func (b *Buffer) Clear() {
	b.bufferSize = 0
	b.buffer = b.buffer[:0]
}
