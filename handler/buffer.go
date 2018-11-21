package handler

import (
	"github.com/syyongx/llog/types"
)

type Buffer struct {
	Handler

	handler types.IHandler
	records chan *types.Record
	close   chan bool
}

// New async handler
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

// Handles a record.
func (b *Buffer) Handle(record *types.Record) bool {
	if !b.IsHandling(record) {
		return false
	}

	b.records <- record

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
	b.records <- nil
	<-b.close
}
