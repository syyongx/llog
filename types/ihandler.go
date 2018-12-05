package types

// IHandler Interface of the Handler.
type IHandler interface {
	// Checks whether the given record will be handled by this handler.
	IsHandling(record *Record) bool

	// Handles a record.
	Handle(record *Record) bool

	// Handles a set of records at once.
	HandleBatch(records []*Record)

	// Closes the handler.
	Close()
}

// FormattableHandler Interface to describe loggers that have a formatter
type FormattableHandler interface {
	// Sets the formatter.
	SetFormatter(formatter *Formatter) *IHandler

	// Gets the formatter.
	GetFormatter() *Formatter
}

// ProcessableHandler Interface to describe loggers that have processors
type ProcessableHandler interface {
	// Adds a processor in the stack.
	PushProcessor() *IHandler

	// Removes the processor on top of the stack and returns it.
	PopProcessor()
}
