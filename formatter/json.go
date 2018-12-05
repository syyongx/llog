package formatter

import (
	"github.com/syyongx/llog/types"
)

// DefaultFields for a log record
var DefaultFields = []string{
	"Datetime",
	"Channel",
	"LevelName",
	"Message",
	"Context",
	"Extra",
}

// JSON struct definition
type JSON struct {
	Normalizer

	fileds        []string
	appendNewline bool
}

// NewJSON appendNewline: Is append new line.
func NewJSON(fields []string, appendNewline bool) *JSON {
	if fields == nil {
		fields = DefaultFields
	}
	j := &JSON{
		fileds:        fields,
		appendNewline: appendNewline,
	}
	return j
}

// IsAppendNewLine is append new line
func (j *JSON) IsAppendNewLine() bool {
	return j.appendNewline
}

// Format a record
func (j *JSON) Format(record *types.Record) error {
	output := make(map[string]string, len(j.fileds))
	for _, field := range j.fileds {
		switch field {
		case "Datetime":
			output[field] = j.normalizeTime(record.Datetime)
		case "Channel":
			output[field] = record.Channel
		case "LevelName":
			output[field] = record.LevelName
		case "Message":
			output[field] = record.Message
		case "Context":
			output[field] = j.normalizeContext(record.Context)
		case "Extra":
			output[field] = j.normalizeExtra(record.Extra)
		default:
			output[field] = "unknow"
		}
	}
	record.Formatted.Write(j.JSON(output))
	if j.appendNewline {
		record.Formatted.WriteRune('\n')
	}
	return nil
}

// FormatBatch Format batch record
func (j *JSON) FormatBatch(records []*types.Record) error {
	for _, record := range records {
		err := j.Format(record)
		if err != nil {
			return err
		}
	}
	return nil
}
