package handler

import (
	"github.com/syyongx/llog/types"
)

type Buffer struct {
	Handler
	Processable

	handler types.IHandler
	buffer  chan *types.Record
	close   chan bool
}

// New buffer handler
// bufferSize: buffer size.
func NewBuffer(handler types.IHandler, bufferSize, level int, bubble bool) *Buffer {
	buf := &Buffer{
		handler: handler,
		buffer:  make(chan *types.Record, bufferSize),
		close:   make(chan bool, 0),
	}
	buf.SetLevel(level)
	buf.SetBubble(bubble)

	go func() {
		for {
			record := <-buf.buffer
			if record == nil {
				break
			}
			buf.handler.Handle(record)
		}
	}()

	return buf
}

// Handles a record.
func (b *Buffer) Handle(record *types.Record) bool {
	if record == nil {
		b.handler.Close()
		b.close <- true
		return false
	}
	if !b.IsHandling(record) {
		return false
	}

	b.buffer <- record

	return false == b.GetBubble()
}

// Handles a set of records.
func (b *Buffer) HandleBatch(records []*types.Record) {
	for _, record := range records {
		b.Handle(record)
	}
}

// close
func (b *Buffer) Close() {
	b.buffer <- nil
	<-b.close
}
