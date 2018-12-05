package formatter

import (
	"encoding/json"
	"github.com/syyongx/llog/types"
	"math"
	"strconv"
	"time"
)

// Normalize incoming records to remove objects/resources so it's easier to dump to various targets
type Normalizer struct {
	dateFormat string
}

// NewNormalizer New normalizer
func NewNormalizer(dateFormat string) *Normalizer {
	return &Normalizer{
		dateFormat: dateFormat,
	}
}

// SetDateFormat Set dateFormat
func (n *Normalizer) SetDateFormat(dateFormat string) {
	n.dateFormat = dateFormat
}

// DateFormat Get dateFormat
func (n *Normalizer) DateFormat() string {
	return n.dateFormat
}

// Normalize extra of record
func (n *Normalizer) normalizeExtra(extra types.RecordExtra) string {
	if len(extra) > 1000 {
		i := len(extra) - 1000
		for k := range extra {
			if i--; i < 0 {
				break
			}
			delete(extra, k)
		}
	}
	// fmt.Sprintf("Over 1000 items (%d total), aborting normalization", len(data.(types.RecordExtra)));
	return string(n.Json(extra))
}

// Normalize context of record
func (n *Normalizer) normalizeContext(ctx types.RecordContext) string {
	if len(ctx) > 1000 {
		i := len(ctx) - 1000
		for k := range ctx {
			if i--; i < 0 {
				break
			}
			delete(ctx, k)
		}
	}
	return string(n.Json(ctx))
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
		}

		return "-INF"
	}

	if math.IsNaN(f) {
		return "NaN"
	}
	return strconv.FormatFloat(f, 'f', 3, 64)
}

// Json Return the JSON representation of a value
func (n *Normalizer) Json(data interface{}) []byte {
	v, err := json.Marshal(data)
	if err != nil {
		return []byte(err.Error())
	}
	return v
}
