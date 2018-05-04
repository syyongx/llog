package formatter

import "github.com/syyongx/llog/types"

// interface fromatter
type IFormatter interface {
	// Formats a log record.
	Format(record types.Record) ([]byte, error)

	// Formats a set of log records.
	FormatBatch(records []types.Record) ([]byte, error)
}
