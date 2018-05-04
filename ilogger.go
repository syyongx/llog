package llog

// Logger Interface
// RFC 5424
type ILogger interface {
	// System is unusable.
	Emergency(message interface{})

	// Action must be taken immediately.
	// Example: Entire website down, database unavailable, etc. This should
	// trigger the SMS alerts and wake you up.
	Alert(message interface{})

	// Critical conditions.
	// Example: Application component unavailable, unexpected exception.
	Critical(message interface{})

	// Runtime errors that do not require immediate action but should typically
	// be logged and monitored.
	Error(message interface{})

	// Exceptional occurrences that are not errors.
	// Example: Use of deprecated APIs, poor use of an API, undesirable things
	// that are not necessarily wrong.
	Warning(message interface{})

	// Normal but significant events.
	Notice(message interface{})

	// Interesting events.
	// Example: User logs in, SQL logs.
	Info(message interface{})

	// Detailed debug information.
	Debug(message interface{})

	// Logs with an arbitrary level.
	Log(level int, message interface{})
}
