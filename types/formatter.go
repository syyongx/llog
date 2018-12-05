package types

// Formatter interface fromatter
type Formatter interface {
	// Format formats a log record.
	Format(record *Record) error

	// FormatBatch formats a set of log records.
	FormatBatch(records []*Record) error
}
