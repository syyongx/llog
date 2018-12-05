package types

import (
	"bytes"
	"sync"
	"time"
)

var recordPool *sync.Pool

func init() {
	recordPool = &sync.Pool{
		New: func() interface{} {
			record := new(Record)
			record.Formatted = new(bytes.Buffer)
			return record
		},
	}
}

// RecordContext record context map
type RecordContext map[string]interface{}
// RecordExtra record extra map
type RecordExtra map[string]interface{}

// Record struct definition
type Record struct {
	Level     int
	LevelName string
	Channel   string
	Datetime  time.Time
	Message   string
	Context   RecordContext
	Extra     RecordExtra
	Formatted *bytes.Buffer
}

// NewRecord Get record from pool.
func NewRecord() *Record {
	record := new(Record)
	record.Formatted = new(bytes.Buffer)
	return record
}

// GetRecord Get record
func GetRecord() *Record {
	record, ok := recordPool.Get().(*Record)
	if !ok {
		return NewRecord()
	}
	return record
}

// ReleaseRecord Put record to pool.
func ReleaseRecord(record *Record) {
	if record == nil {
		return
	}
	record.Formatted.Reset()
	recordPool.Put(record)
}
