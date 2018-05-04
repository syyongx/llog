package handler

import (
	"github.com/syyongx/llog/types"
	"os"
)

type File struct {
	Handler
	Processable
	Formattable

	useLocking bool
	writer     *os.File
}

// New file.
func NewFile(path string, level int, bubble bool, useLocking bool) (IHandler, error) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	f := &File{
		useLocking: useLocking,
		writer:     fd,
	}
	f.SetLevel(level)
	f.SetBubble(bubble)
	return f, nil
}

// Handle
func (f *File) Handle(record types.Record) bool {
	if !f.IsHandling(record) {
		return false
	}
	if f.processors != nil {
		record = f.ProcessRecord(record)
	}
	var err error
	record["formatted"], err = f.GetFormatter().Format(record)
	if err != nil {
		return false
	}

	f.Write(record)

	return false == f.GetBubble()
}

// HandleBatch
func (f *File) HandleBatch(records []types.Record) {
	for _, record := range records {
		f.Handle(record)
	}
}

// Write to file.
func (f *File) Write(record types.Record) {
	f.writer.Write(record["formatted"].([]byte))
	defer f.Close()
}

// Close wreter
func (f *File) Close() {
	f.writer.Close()
	f.writer = nil
}
