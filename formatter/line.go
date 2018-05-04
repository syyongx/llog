package formatter

import (
	"github.com/syyongx/llog/types"
	"strings"
	"fmt"
)

// Formats incoming records into a one-line string
// This is especially useful for logging to files
type Line struct {
	Normalizer

	format string
}

func NewLine(format, dateFormat string) IFormatter {
	l := &Line{
		format: format,
	}
	l.SetDateFormat(dateFormat)
	return l
}

func (l *Line) Format(record types.Record) ([]byte, error) {
	output := l.format

	vars := l.Normalize(record, 0).(types.Record)
	for k, v := range vars {
		placeholder := "%" + k + "%"
		if strings.Contains(output, placeholder) {
			output = strings.Replace(output, placeholder, l.String(v), 1)
		}
	}

	return []byte(output), nil
}

//
func (l *Line) FormatBatch(records []types.Record) ([]byte, error) {
	var message []byte
	for _, record := range records {
		entry, err := l.Format(record)
		if err != nil {
			return message, err
		}
		message = append(message, entry...)
	}
	return message, nil
}

// stringfy
func (l *Line) String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128:
		return fmt.Sprintf("%v", data)
	}
	return l.ToJson(data)
}
