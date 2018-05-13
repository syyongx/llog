package types

import (
	"time"
	"bytes"
	"sync"
)

var recordPool *sync.Pool

func init() {
	recordPool = &sync.Pool{
		New: func() interface{} {
			return new(Record)
		},
	}
}

type RecordContext map[string]interface{}
type RecordExtra map[string]interface{}

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

// Get record from pool.
func NewRecord() *Record {
	record, ok := recordPool.Get().(*Record);
	if !ok {
		record = new(Record)
		//record.Buffer = new(bytes.Buffer)
	}
	return record
}

// Put record to pool.
func ReleaseRecord(record *Record) {
	record.Formatted.Reset()
	recordPool.Put(record)
}
