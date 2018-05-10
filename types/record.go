package types

import (
	"time"
	"bytes"
	"sync"
)

var BufferPool *sync.Pool
var RecordPool *sync.Pool

func init() {
	BufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	RecordPool = &sync.Pool{
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
	Buffer    *bytes.Buffer
}

// Get record from pool.
func GetRecord() *Record {
	if record, ok := RecordPool.Get().(*Record); ok {
		return record
	}
	return new(Record)
}

// Put record to pool.
func PutRecord(record *Record) {
	RecordPool.Put(record)
}
