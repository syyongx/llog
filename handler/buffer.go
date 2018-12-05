package handler

import (
	"github.com/syyongx/llog/types"
)

// Buffer struct definition
type Buffer struct {
	Handler

	handler types.IHandler
	records chan *types.Record
	close   chan bool
}

// NewBuffer New async handler
// bufSize: channel buffer size.
func NewBuffer(handler types.IHandler, bufSize, level int, bubble bool) *Buffer {
	buf := &Buffer{
		handler: handler,
		records: make(chan *types.Record, bufSize),
		close:   make(chan bool, 0),
	}
	buf.SetLevel(level)
	buf.SetBubble(bubble)

	go func() {
		for {
			record := <-buf.records
			if record == nil {
				buf.handler.Close()
				buf.close <- true
				break
			}
			buf.handler.Handle(record)
		}
	}()

	return buf
}

// Handle handles a record.
func (b *Buffer) Handle(record *types.Record) bool {
	if !b.IsHandling(record) {
		return false
	}

	b.records <- record

	return false == b.GetBubble()
}

// HandleBatch Handles a set of records.
func (b *Buffer) HandleBatch(records []*types.Record) {
	for _, record := range records {
		b.Handle(record)
	}
}

// Close close
func (b *Buffer) Close() {
	b.records <- nil
	<-b.close
}
