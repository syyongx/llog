package formatter

import (
	"math"
	"github.com/syyongx/llog/types"
	"encoding/json"
	"time"
	"strconv"
)

// Normalizes incoming records to remove objects/resources so it's easier to dump to various targets
type Normalizer struct {
	dateFormat string
}

func NewNormalizer(dateFormat string) *Normalizer {
	return &Normalizer{
		dateFormat: dateFormat,
	}
}

// Set dateFormat
func (n *Normalizer) SetDateFormat(dateFormat string) {
	n.dateFormat = dateFormat
}

// Get dateFormat
func (n *Normalizer) GetDateFormat() string {
	return n.dateFormat
}

// Normalize extra of record
func (n *Normalizer) normalizeExtra(extra types.RecordExtra) string {
	if len(extra) > 1000 {
		return ""
	}
	// fmt.Sprintf("Over 1000 items (%d total), aborting normalization", len(data.(types.RecordExtra)));
	return n.ToJson(extra)
}

// Normalize context of record
func (n *Normalizer) normalizeContext(ctx types.RecordContext) string {
	if len(ctx) > 1000 {
		return ""
	}
	return n.ToJson(ctx)
}

// Normalize float
func (n *Normalizer) normalizeTime(t time.Time) string {
	return t.Format(n.dateFormat)
}

// Normalize float
func (n *Normalizer) normalizeFloat(f float64) string {
	// whether n is an infinity
	if math.IsInf(f, 0) {
		if f > 0 {
			return "INF"
		} else {
			return "-INF"
		}
	}
	if math.IsNaN(f) {
		return "NaN"
	}
	return strconv.FormatFloat(f, 'G', 30, 64)
}

// Return the JSON representation of a value
func (n *Normalizer) ToJson(data interface{}) string {
	v, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(v)
}
