package types

import (
	"time"
	"bytes"
	"sync"
)

var BufferPool *sync.Pool

func init() {
	BufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
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

func NewRecord() *Record {
	return &Record{}
}
