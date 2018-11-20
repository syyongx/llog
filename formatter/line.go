package formatter

import (
	"fmt"
	"github.com/syyongx/llog/types"
	"strings"
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
	oldnew := []string{
		"%Datetime%", l.normalizeTime(record.Datetime),
		"%Channel%", record.Channel,
		"%LevelName%", record.LevelName,
		"%Message%", record.Message,
		"%Context%", l.normalizeContext(record.Context),
		"%Extra%", l.normalizeExtra(record.Extra),
	}
	output := strings.NewReplacer(oldnew...).Replace(l.format)
	_, err := record.Formatted.WriteString(output)
	return err
}

// Batch format records.
func (l Line) FormatBatch(records []*types.Record) error {
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
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		return fmt.Sprintf("%v", data)
	}
	return string(l.Json(data))
}
