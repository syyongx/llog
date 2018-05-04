package formatter

import (
	"fmt"
	"math"
	"github.com/syyongx/llog/types"
	"encoding/json"
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

//func (n *Normalizer) Format(record types.Record) ([]byte, error) {
//	return n.Normalize(record, 0)
//}

func (n *Normalizer) Normalize(data interface{}, depth int) interface{} {
	if depth > 9 {
		return "Over 9 levels deep, aborting normalization"
	}
	switch data.(type) {
	case float32, float64:
		n := float64(data.(float32))
		// whether n is an infinity
		if math.IsInf(n, 0) {
			if n > 0 {
				return "INF"
			} else {
				return "-INF"
			}
		}
		if math.IsNaN(n) {
			return "NaN"
		}
		return data
	case types.Record:
		normalized := make(types.Record)
		count := 1
		for k, v := range data.(types.Record) {
			if count ++; count >= 1000 {
				normalized["..."] = fmt.Sprintf("Over 1000 items (%d total), aborting normalization", len(data.(types.Record)));
				break
			}
			normalized[k] = n.Normalize(v, depth+1)
		}
		return normalized
	}

	return data
}

// Return the JSON representation of a value
func (n *Normalizer) ToJson(data interface{}) string {
	v, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(v)
}
