package handler

import (
	"os"
	"github.com/syyongx/llog/types"
)

// File handler.
type File struct {
	Handler
	Processable
	Formattable

	Fd *os.File
}

// New file handler
func NewFile(path string, level int, bubble bool, filePerm os.FileMode) (*File, error) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, filePerm)
	if err != nil {
		return nil, err
	}
	f := &File{
		Fd: fd,
	}
	f.SetLevel(level)
	f.SetBubble(bubble)
	return f, nil
}

// Handle
func (f *File) Handle(record *types.Record) bool {
	if !f.IsHandling(record) {
		return false
	}
	if f.processors != nil {
		f.ProcessRecord(record)
	}
	err := f.GetFormatter().Format(record)
	if err != nil {
		return false
	}
	f.Write(record)

	return false == f.GetBubble()
}

// HandleBatch
func (f *File) HandleBatch(records []*types.Record) {
	for _, record := range records {
		f.Handle(record)
	}
}

// Write to file.
func (f *File) Write(record *types.Record) {
	if f.Fd == nil {
		return
	}
	f.Fd.Write(record.Formatted.Bytes())
	//defer f.Close()
}

// Close writer
func (f *File) Close() {
	f.Fd.Close()
	f.Fd = nil
}
