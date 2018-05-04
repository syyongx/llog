package types

import "time"

type RecordContext map[string]interface{}
type RecordExtra map[string]interface{}

type Record struct {
	Message   string
	Level     int
	LevelName string
	Channel   string
	Datetime  time.Time
	Context   RecordContext
	Extra     RecordExtra
	Formatted []byte
}

func NewRecord() *Record {
	return &Record{}
}
