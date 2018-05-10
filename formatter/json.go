package formatter

import (
	"github.com/syyongx/llog/types"
)

var DefaultFields = []string{
	"Datetime",
	"Channel",
	"LevelName",
	"Message",
	"Context",
	"Extra",
}

type Json struct {
	Normalizer

	fileds        []string
	appendNewline bool
}

// appendNewline: Is append new line.
func NewJson(fields []string, appendNewline bool) *Json {
	if fields == nil {
		fields = DefaultFields
	}
	j := &Json{
		fileds:        fields,
		appendNewline: appendNewline,
	}
	return j
}

func (j *Json) IsAppendNewLine() bool {
	return j.appendNewline
}

// Format a record
func (j *Json) Format(record *types.Record) error {
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
	record.Buffer.Write(j.ToJson(output))
	if j.appendNewline {
		record.Buffer.WriteRune('\n')
	}
	return nil
}

// Format batch record
func (j *Json) FormatBatch(records []*types.Record) error {
	for _, record := range records {
		err := j.Format(record)
		if err != nil {
			return err
		}
	}
	return nil
}
