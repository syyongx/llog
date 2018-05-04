package handler

import (
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/types"
)

// Interface of the Handler.
type IHandler interface {
	// Checks whether the given record will be handled by this handler.
	IsHandling(record *types.Record) bool

	// Handles a record.
	Handle(record *types.Record) bool

	// Handles a set of records at once.
	HandleBatch(records []*types.Record)

	// Closes the handler.
	Close()
}

// Interface to describe loggers that have a formatter
type IFormattableHandler interface {
	// Sets the formatter.
	SetFormatter(formatter *formatter.IFormatter) *IHandler

	// Gets the formatter.
	GetFormatter() *formatter.IFormatter
}

// Interface to describe loggers that have processors
type IProcessableHandler interface {
	// Adds a processor in the stack.
	PushProcessor() *IHandler

	// Removes the processor on top of the stack and returns it.
	PopProcessor()
}
