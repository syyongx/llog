package formatter

import (
	"github.com/syyongx/llog/types"
	"strings"
	"fmt"
)

var DefaultFormat = "[%datetime%] %channel%.%levelName%: %message% %context% %extra%\n"

// Formats incoming records into a one-line string
// This is especially useful for logging to files
type Line struct {
	Normalizer

	format string
}

func NewLine(format, dateFormat string) *Line {
	if format == "" {
		format = DefaultFormat
	}
	l := &Line{
		format: format,
	}
	l.SetDateFormat(dateFormat)
	return l
}

func (l *Line) Format(record *types.Record) error {
	output := l.format

	if strings.Contains(l.format, "%datetime%") {
		output = strings.Replace(output, "%datetime%", l.normalizeTime(record.Datetime), 1)
	}
	if strings.Contains(l.format, "%channel%") {
		output = strings.Replace(output, "%channel%", record.Channel, 1)
	}
	if strings.Contains(l.format, "%levelName%") {
		output = strings.Replace(output, "%levelName%", record.LevelName, 1)
	}
	if strings.Contains(l.format, "%message%") {
		output = strings.Replace(output, "%message%", record.Message, 1)
	}
	if strings.Contains(l.format, "%context%") {
		output = strings.Replace(output, "%context%", l.normalizeContext(record.Context), 1)
	}
	if strings.Contains(l.format, "%extra%") {
		output = strings.Replace(output, "%extra%", l.normalizeExtra(record.Extra), 1)
	}
	record.Buffer.WriteString(output)

	return nil
}

// Batch format records.
func (l *Line) FormatBatch(records []*types.Record) error {
	for _, record := range records {
		err := l.Format(record)
		if err != nil {
			return err
		}
	}
	return nil
}

// stringfy
func (l *Line) String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128:
		return fmt.Sprintf("%v", data)
	}
	return string(l.ToJson(data))
}
