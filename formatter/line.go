package formatter

import (
	"github.com/syyongx/llog/types"
	"strings"
	"fmt"
)

var DefaultFormat = "[%Datetime%] %Channel%.%LevelName%: %Message% %Context% %Extra%\n"

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

	if strings.Contains(l.format, "%Datetime%") {
		output = strings.Replace(output, "%Datetime%", l.normalizeTime(record.Datetime), 1)
	}
	if strings.Contains(l.format, "%Channel%") {
		output = strings.Replace(output, "%Channel%", record.Channel, 1)
	}
	if strings.Contains(l.format, "%LevelName%") {
		output = strings.Replace(output, "%LevelName%", record.LevelName, 1)
	}
	if strings.Contains(l.format, "%Message%") {
		output = strings.Replace(output, "%Message%", record.Message, 1)
	}
	if strings.Contains(l.format, "%Context%") {
		output = strings.Replace(output, "%Context%", l.normalizeContext(record.Context), 1)
	}
	if strings.Contains(l.format, "%Extra%") {
		output = strings.Replace(output, "%Extra%", l.normalizeExtra(record.Extra), 1)
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
